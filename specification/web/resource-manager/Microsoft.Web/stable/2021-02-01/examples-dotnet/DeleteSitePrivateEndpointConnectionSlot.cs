using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AppService;
using Azure.ResourceManager.AppService.Models;

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2021-02-01/examples/DeleteSitePrivateEndpointConnectionSlot.json
// this example is just showing the usage of "WebApps_DeletePrivateEndpointConnectionSlot" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SiteSlotPrivateEndpointConnectionResource created on azure
// for more information of creating SiteSlotPrivateEndpointConnectionResource, please refer to the document of SiteSlotPrivateEndpointConnectionResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "rg";
string name = "testSite";
string slot = "stage";
string privateEndpointConnectionName = "connection";
ResourceIdentifier siteSlotPrivateEndpointConnectionResourceId = SiteSlotPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name, slot, privateEndpointConnectionName);
SiteSlotPrivateEndpointConnectionResource siteSlotPrivateEndpointConnection = client.GetSiteSlotPrivateEndpointConnectionResource(siteSlotPrivateEndpointConnectionResourceId);

// invoke the operation
ArmOperation<BinaryData> lro = await siteSlotPrivateEndpointConnection.DeleteAsync(WaitUntil.Completed);
BinaryData result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
