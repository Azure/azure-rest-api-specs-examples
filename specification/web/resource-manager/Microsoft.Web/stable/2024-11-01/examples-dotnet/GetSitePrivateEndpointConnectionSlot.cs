using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService.Models;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/GetSitePrivateEndpointConnectionSlot.json
// this example is just showing the usage of "WebApps_GetPrivateEndpointConnectionSlot" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this WebSiteSlotResource created on azure
// for more information of creating WebSiteSlotResource, please refer to the document of WebSiteSlotResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "rg";
string name = "testSite";
string slot = "stage";
ResourceIdentifier webSiteSlotResourceId = WebSiteSlotResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name, slot);
WebSiteSlotResource webSiteSlot = client.GetWebSiteSlotResource(webSiteSlotResourceId);

// get the collection of this SiteSlotPrivateEndpointConnectionResource
SiteSlotPrivateEndpointConnectionCollection collection = webSiteSlot.GetSiteSlotPrivateEndpointConnections();

// invoke the operation
string privateEndpointConnectionName = "connection";
NullableResponse<SiteSlotPrivateEndpointConnectionResource> response = await collection.GetIfExistsAsync(privateEndpointConnectionName);
SiteSlotPrivateEndpointConnectionResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    RemotePrivateEndpointConnectionARMResourceData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
