// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This represents a container app component resource
 */
export class Containerapp extends pulumi.ComponentResource {
    /** @internal */
    public static readonly __pulumiType = 'azure-justrun:index:containerapp';

    /**
     * Returns true if the given object is an instance of Containerapp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Containerapp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Containerapp.__pulumiType;
    }

    /**
     * The URL of the container app
     */
    public /*out*/ readonly url!: pulumi.Output<string | undefined>;

    /**
     * Create a Containerapp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ContainerappArgs, opts?: pulumi.ComponentResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["dockerImageName"] = args ? args.dockerImageName : undefined;
            resourceInputs["imageDirectory"] = args ? args.imageDirectory : undefined;
            resourceInputs["namePrefix"] = args ? args.namePrefix : undefined;
            resourceInputs["registryName"] = args ? args.registryName : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["storageAccountName"] = args ? args.storageAccountName : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["url"] = undefined /*out*/;
        } else {
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Containerapp.__pulumiType, name, resourceInputs, opts, true /*remote*/);
    }
}

/**
 * The set of arguments for constructing a Containerapp resource.
 */
export interface ContainerappArgs {
    /**
     * The name of the docker image. One will be created if not provided
     */
    dockerImageName?: pulumi.Input<string>;
    /**
     * The name of the directory where the docker image to be created is. NOT the actual directory, i.e. 'nodeapp' instead of './nodeapp'
     */
    imageDirectory?: pulumi.Input<string>;
    /**
     * The name prefix given to child resources of this component. Should not contain dashes.
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * The name of the image registry. Must belong to the resource group specified in ResourceGroupName. One will be created if not provided.
     */
    registryName?: pulumi.Input<string>;
    /**
     * The resource group to use. One will be created if not provided.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * The name of the storage account to use. One will be created if not provided.
     */
    storageAccountName?: pulumi.Input<string>;
    /**
     * The version of the created docker image
     */
    version?: pulumi.Input<string>;
}
