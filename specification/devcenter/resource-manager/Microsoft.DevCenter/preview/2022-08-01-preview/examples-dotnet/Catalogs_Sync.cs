using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DevCenter;
using Azure.ResourceManager.DevCenter.Models;

// Generated from example definition: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/Catalogs_Sync.json
// this example is just showing the usage of "Catalogs_Sync" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CatalogResource created on azure
// for more information of creating CatalogResource, please refer to the document of CatalogResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "rg1";
string devCenterName = "Contoso";
string catalogName = "{catalogName}";
ResourceIdentifier catalogResourceId = CatalogResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, devCenterName, catalogName);
CatalogResource catalog = client.GetCatalogResource(catalogResourceId);

// invoke the operation
await catalog.SyncAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
