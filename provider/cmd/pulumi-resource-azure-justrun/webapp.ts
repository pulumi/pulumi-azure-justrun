import * as pulumi from "@pulumi/pulumi";
import * as storage from "@pulumi/azure-native/storage";
import * as resources from "@pulumi/azure-native/resources";
import * as web from "@pulumi/azure-native/web";

export class WebApp extends pulumi.ComponentResource {

    public readonly url!: pulumi.Output<string>;

    constructor(name: string, args: WebAppArgs, opts: pulumi.ResourceOptions = {}) {
        super("azure-justrun:index:webapp", name, {}, opts);
        const namePrefix = args.namePrefix ?? ""

        const resourceGroup = args.resourceGroup ?? new resources.ResourceGroup(`${namePrefix}rg`, {}, {parent: this});

        // Storage Account name must be lowercase and cannot have any dash characters
        const storageAccount = args.storageAccount ?? new storage.StorageAccount(`${namePrefix}sa`, {
            resourceGroupName: resourceGroup.name,
            kind: storage.Kind.StorageV2,
            sku: {
                name: args.storageSkuName ?? storage.SkuName.Standard_LRS,
            },
        }, {parent: this});


        const appServicePlan = new web.AppServicePlan(`${namePrefix}asp`, {
            resourceGroupName: resourceGroup.name,
            kind: "App",
            sku: {
                name: args.appSkuName ?? "B1",
                tier: args.appSkuTier ?? "Basic",
            },
        }, {parent: this});

        const storageContainer = new storage.BlobContainer(`${namePrefix}container`, {
            resourceGroupName: resourceGroup.name,
            accountName: storageAccount.name,
            publicAccess: args.containerPublicAccess ?? storage.PublicAccess.None,
        }, {parent: this});

        const blob = new storage.Blob(`${namePrefix}blob`, {
            resourceGroupName: resourceGroup.name,
            accountName: storageAccount.name,
            containerName: storageContainer.name,
            source: new pulumi.asset.FileArchive(args.filePath ?? "../wwwroot"),
        }, {parent: this});

        const codeBlobUrl = pulumi.all(
            [storageAccount.name, storageContainer.name, blob.name, resourceGroup.name]).apply(
            (args: any) => getSASToken(args[0], args[1], args[2], args[3]));

        const app = new web.WebApp(`${namePrefix}webapp`, {
            resourceGroupName: resourceGroup.name,
            serverFarmId: appServicePlan.id,
            siteConfig: {
                appSettings: [
                    {
                        name: "WEBSITE_RUN_FROM_PACKAGE",
                        value: codeBlobUrl,
                    },
                ],
            },
        }, {parent: this});
        

        this.url = pulumi.interpolate `https://${app.defaultHostName}`;
        this.registerOutputs();
    }   
}

function getSASToken(storageAccountName: string, storageContainerName: string, blobName: string, resourceGroupName: string) {
    const blobSAS = storage.listStorageAccountServiceSAS({
        accountName: storageAccountName,
        protocols: storage.HttpProtocol.Https,
        sharedAccessStartTime: "2021-01-01",
        sharedAccessExpiryTime: "2030-01-01",
        resource: storage.SignedResource.C,
        resourceGroupName: resourceGroupName,
        permissions: storage.Permissions.R,
        canonicalizedResource: "/blob/" + storageAccountName + "/" + storageContainerName,
        contentType: "application/json",
        cacheControl: "max-age=5",
        contentDisposition: "inline",
        contentEncoding: "deflate",
    });
    return pulumi.interpolate `https://${storageAccountName}.blob.core.windows.net/${storageContainerName}/${blobName}?${blobSAS.then((x: any) => x.serviceSasToken)}`;
}


export interface WebAppArgs extends pulumi.ComponentResourceOptions{
    appSkuName?: string,
    appSkuTier?: string,
    filePath?: string,
    containerPublicAccess?: storage.PublicAccess,
    storageSkuName?: storage.SkuName,
    storageAccount?: storage.StorageAccount,
    resourceGroup?: resources.ResourceGroup,
    namePrefix?: string
}

