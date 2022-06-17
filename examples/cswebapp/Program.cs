using Pulumi;
using Pulumi.AzureJustrun;
using System.Collections.Generic;

await Pulumi.Deployment.RunAsync(() =>
{
    // Create an Azure Resource Group
    var resourceGroup = new Webapp("myapp", new WebappArgs{
        FilePath = "./www"
    });

    // Export the primary key of the Storage Account
    return new Dictionary<string, object?>
    {
        ["url"] = resourceGroup.Url
    };
});
