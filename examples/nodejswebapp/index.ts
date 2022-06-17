import * as pulumi from "@pulumi/pulumi";
import * as justrun from "@pulumi/azure-justrun"
import * as resources from "@pulumi/azure-native/resources";
import * as storage from "@pulumi/azure-native/storage";

// Create an Azure Resource Group
const rg = new resources.ResourceGroup("rg", {
})

const storageAccount = new storage.StorageAccount("sa", {
    resourceGroupName: rg.name,
    kind: storage.Kind.StorageV2,
    sku: {
        name: storage.SkuName.Standard_LRS,
    },
});

const webapp = new justrun.Webapp("webapp", {
    "filePath": "./www",
    "storageAccount": storageAccount,
    "resourceGroup": rg,
    "namePrefix": "pref"
})

export const url = webapp.url