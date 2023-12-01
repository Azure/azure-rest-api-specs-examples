using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AppService;
using Azure.ResourceManager.AppService.Models;

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2021-02-01/examples/ApproveRejectSitePrivateEndpointConnection.json
// this example is just showing the usage of "WebApps_ApproveOrRejectPrivateEndpointConnection" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SitePrivateEndpointConnectionResource created on azure
// for more information of creating SitePrivateEndpointConnectionResource, please refer to the document of SitePrivateEndpointConnectionResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "rg";
string name = "testSite";
string privateEndpointConnectionName = "connection";
ResourceIdentifier sitePrivateEndpointConnectionResourceId = SitePrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name, privateEndpointConnectionName);
SitePrivateEndpointConnectionResource sitePrivateEndpointConnection = client.GetSitePrivateEndpointConnectionResource(sitePrivateEndpointConnectionResourceId);

// invoke the operation
PrivateLinkConnectionApprovalRequestInfo info = new PrivateLinkConnectionApprovalRequestInfo()
{
    PrivateLinkServiceConnectionState = new PrivateLinkConnectionState()
    {
        Status = "Approved",
        Description = "Approved by admin.",
        ActionsRequired = "",
    },
};
ArmOperation<SitePrivateEndpointConnectionResource> lro = await sitePrivateEndpointConnection.UpdateAsync(WaitUntil.Completed, info);
SitePrivateEndpointConnectionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
RemotePrivateEndpointConnectionARMResourceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
