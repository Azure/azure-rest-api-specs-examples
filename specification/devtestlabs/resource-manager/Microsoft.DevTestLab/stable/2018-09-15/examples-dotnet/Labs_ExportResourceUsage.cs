using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DevTestLabs;
using Azure.ResourceManager.DevTestLabs.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Resources.Models;

// Generated from example definition: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Labs_ExportResourceUsage.json
// this example is just showing the usage of "Labs_ExportResourceUsage" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DevTestLabResource created on azure
// for more information of creating DevTestLabResource, please refer to the document of DevTestLabResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "resourceGroupName";
string name = "{labName}";
ResourceIdentifier devTestLabResourceId = DevTestLabResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name);
DevTestLabResource devTestLab = client.GetDevTestLabResource(devTestLabResourceId);

// invoke the operation
DevTestLabExportResourceUsageContent content = new DevTestLabExportResourceUsageContent()
{
    BlobStorageAbsoluteSasUri = new Uri("https://invalid.blob.core.windows.net/export.blob?sv=2015-07-08&sig={sas}&sp=rcw"),
    UsageStartOn = DateTimeOffset.Parse("2020-12-01T00:00:00Z"),
};
await devTestLab.ExportResourceUsageAsync(WaitUntil.Completed, content);

Console.WriteLine($"Succeeded");
