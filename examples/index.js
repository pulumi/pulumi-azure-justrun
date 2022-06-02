"use strict";
const pulumi = require("@pulumi/pulumi");
const azure = require("@pulumi/azure-native");
const azwebex = require("../sdk/nodejs");

// Create an Azure Resource Group
const webapp = new azwebex.index.StaticWebsite()
// Export the primary storage key for the storage account
exports.url = webapp.url
