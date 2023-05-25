using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DevCenter;
using Azure.ResourceManager.DevCenter.Models;

// Generated from example definition: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2023-04-01/examples/Catalogs_Get.json
// this example is just showing the usage of "Catalogs_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DevCenterResource created on azure
// for more information of creating DevCenterResource, please refer to the document of DevCenterResource
string subscriptionId = "0ac520ee-14c0-480f-b6c9-0a90c58ffff";
string resourceGroupName = "rg1";
string devCenterName = "Contoso";
ResourceIdentifier devCenterResourceId = DevCenterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, devCenterName);
DevCenterResource devCenter = client.GetDevCenterResource(devCenterResourceId);

// get the collection of this DevCenterCatalogResource
DevCenterCatalogCollection collection = devCenter.GetDevCenterCatalogs();

// invoke the operation
string catalogName = "CentralCatalog";
bool result = await collection.ExistsAsync(catalogName);

Console.WriteLine($"Succeeded: {result}");
