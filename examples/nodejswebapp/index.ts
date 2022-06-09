import * as pulumi from "@pulumi/pulumi";
import * as justrun from "@pulumi/azure-native"

// Create an Azure Resource Group
const webapp = new justrun.WebApp("webapp", {
    "filePath": "./www",
})

const url = webapp.url
export url
