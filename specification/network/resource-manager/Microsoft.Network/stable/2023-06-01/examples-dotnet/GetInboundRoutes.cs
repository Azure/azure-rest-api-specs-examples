using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Network;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2023-06-01/examples/GetInboundRoutes.json
// this example is just showing the usage of "VirtualHubs_GetInboundRoutes" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualHubResource created on azure
// for more information of creating VirtualHubResource, please refer to the document of VirtualHubResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string virtualHubName = "virtualHub1";
ResourceIdentifier virtualHubResourceId = VirtualHubResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualHubName);
VirtualHubResource virtualHub = client.GetVirtualHubResource(virtualHubResourceId);

// invoke the operation
VirtualHubInboundRoutesContent content = new VirtualHubInboundRoutesContent()
{
    ResourceUri = new Uri("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRouteGateways/exrGw1/expressRouteConnections/exrConn1"),
    ConnectionType = "ExpressRouteConnection",
};
ArmOperation<EffectiveRouteMapRouteList> lro = await virtualHub.GetVirtualHubInboundRoutesAsync(WaitUntil.Completed, content);
EffectiveRouteMapRouteList result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
