using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MigrationDiscoverySap.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.MigrationDiscoverySap;

// Generated from example definition: specification/workloads/resource-manager/Microsoft.Workloads/SAPDiscoverySites/preview/2023-10-01-preview/examples/SAPDiscoverySites_ImportEntities.json
// this example is just showing the usage of "SapDiscoverySites_ImportEntities" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SapDiscoverySiteResource created on azure
// for more information of creating SapDiscoverySiteResource, please refer to the document of SapDiscoverySiteResource
string subscriptionId = "6d875e77-e412-4d7d-9af4-8895278b4443";
string resourceGroupName = "test-rg";
string sapDiscoverySiteName = "SampleSite";
ResourceIdentifier sapDiscoverySiteResourceId = SapDiscoverySiteResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, sapDiscoverySiteName);
SapDiscoverySiteResource sapDiscoverySite = client.GetSapDiscoverySiteResource(sapDiscoverySiteResourceId);

// invoke the operation
ArmOperation<OperationStatusResult> lro = await sapDiscoverySite.ImportEntitiesAsync(WaitUntil.Completed);
OperationStatusResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
