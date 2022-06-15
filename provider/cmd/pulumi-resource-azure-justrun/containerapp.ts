import * as docker from "@pulumi/docker";
import * as pulumi from "@pulumi/pulumi";

import * as containerregistry from "@pulumi/azure-native/containerregistry";
import * as operationalinsights from "@pulumi/azure-native/operationalinsights";
import * as resources from "@pulumi/azure-native/resources";
import * as app from "@pulumi/azure-native/app";


// Define a component for serving a static website on S3
export class ContainerApp extends pulumi.ComponentResource {

    public readonly url!: pulumi.Output<string>;

    constructor(name: string, args: ContainerAppArgs, opts: pulumi.ResourceOptions = {}) {
        super("azure-justrun:index:containerapp", name, {}, opts); // Register this component with name pulumi:examples:S3Folder
        const namePrefix = args.namePrefix ?? ""

        const version = args.version ?? "v1.0.0";

        const resourceGroup = args.resourceGroup ?? new resources.ResourceGroup(`${namePrefix}rg`, {}, {parent: this});;

        const workspace = new operationalinsights.Workspace("${namePrefix}loganalytics", {
            resourceGroupName: resourceGroup.name,
            sku: {
                name: "PerGB2018",
            },
            retentionInDays: 30,
        }, {parent: this});

        const workspaceSharedKeys = operationalinsights.getSharedKeysOutput({
            resourceGroupName: resourceGroup.name,
            workspaceName: workspace.name,
        });

        const managedEnv = new app.ManagedEnvironment("${namePrefix}env", {
            resourceGroupName: resourceGroup.name,
            appLogsConfiguration: {
                destination: "log-analytics",
                logAnalyticsConfiguration: {
                    customerId: workspace.customerId,
                    sharedKey: workspaceSharedKeys.apply(r => r.primarySharedKey!),
                },
            },
        }, {parent: this});

        const registry = args.registry ?? new containerregistry.Registry("${namePrefix}registry", {
            resourceGroupName: resourceGroup.name,
            sku: {
                name: "Basic",
            },
            adminUserEnabled: true,
        });

        const credentials = containerregistry.listRegistryCredentialsOutput({
            resourceGroupName: resourceGroup.name,
            registryName: registry.name,
        });
        const adminUsername = credentials.apply(c => c.username!);
        const adminPassword = credentials.apply(c => c.passwords![0].value!);
        
        if(args.dockerImageName != null || args.imageDirectory != null) {
            throw new Error("Either dockerImageName or imageDirectory must be specified");
        }

        const imageName = args.dockerImageName ?? new docker.Image("${namePrefix}image", {
            imageName: pulumi.interpolate`${registry.loginServer}/${args.imageDirectory}:v${version}`,
            build: { context: `./${args.imageDirectory}` },
            registry: {
                server: registry.loginServer,
                username: adminUsername,
                password: adminPassword,
            },
        },{parent: this}).imageName;

        const containerApp = new app.ContainerApp("${namePrefix}app", {
            resourceGroupName: resourceGroup.name,
            managedEnvironmentId: managedEnv.id,
            configuration: {
                ingress: {
                    external: true,
                    targetPort: 80,
                },
                registries: [{
                    server: registry.loginServer,
                    username: adminUsername,
                    passwordSecretRef: "pwd",
                }],
                secrets: [{
                    name: "pwd",
                    value: adminPassword,
                }],
            },
            template: {
                containers: [{
                    name: "myapp", //Should this be configurable?
                    image: imageName,
                }],
            },
        }, {parent: this});
        this.url = pulumi.interpolate`https://${containerApp.configuration.apply(c => c?.ingress?.fqdn)}`;
        this.registerOutputs();
     }
}

export interface ContainerAppArgs extends pulumi.ComponentResourceOptions{
    namePrefix?: string;
    resourceGroup?: resources.ResourceGroup;
    registry?: containerregistry.Registry;
    dockerImageName?: string;
    imageDirectory?: string;
    version?: string;
}

