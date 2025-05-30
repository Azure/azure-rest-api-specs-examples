using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService.Models;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2024-04-01/examples/ApproveRejectSitePrivateEndpointConnectionSlot.json
// this example is just showing the usage of "WebApps_ApproveOrRejectPrivateEndpointConnectionSlot" operation, for the dependent resources, they will have to be created separately.

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
RemotePrivateEndpointConnectionARMResourceData data = new RemotePrivateEndpointConnectionARMResourceData()
{
    PrivateLinkServiceConnectionState = new PrivateLinkConnectionState()
    {
        Status = "Approved",
        Description = "Approved by admin.",
        ActionsRequired = "",
    },
};
ArmOperation<SiteSlotPrivateEndpointConnectionResource> lro = await siteSlotPrivateEndpointConnection.UpdateAsync(WaitUntil.Completed, data);
SiteSlotPrivateEndpointConnectionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
RemotePrivateEndpointConnectionARMResourceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
