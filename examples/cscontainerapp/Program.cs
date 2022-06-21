using Pulumi;
using Pulumi.AzureJustrun;
using System.Collections.Generic;

await Pulumi.Deployment.RunAsync(() =>
{
    // Create an Azure Resource Group
    var resourceGroup = new Containerapp("myapp", new ContainerappArgs{
        ImageDirectory = "node-app"
    });

    // Export the primary key of the Storage Account
    return new Dictionary<string, object?>
    {
        ["url"] = resourceGroup.Url
    };
});
