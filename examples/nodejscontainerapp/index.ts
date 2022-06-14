import * as pulumi from "@pulumi/pulumi";
import * as justrun from "@pulumi/azure-justrun"
import * as resources from "@pulumi/azure-native/resources";
import * as storage from "@pulumi/azure-native/storage";

// Create an Azure Resource Group
const containerapp = new justrun.ContainerApp("containerapp", {
    "imageDirectory": "./node-app",
});

export const url = containerapp.url