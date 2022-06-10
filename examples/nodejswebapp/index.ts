import * as pulumi from "@pulumi/pulumi";
import * as justrun from "@pulumi/azure-justrun"

// Create an Azure Resource Group
const webapp = new justrun.Webapp("webapp", {
    "filePath": "./www",
})

export const url = webapp.url