using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/EffectiveRoutesListForConnection.json
// this example is just showing the usage of "VirtualHubs_GetEffectiveVirtualHubRoutes" operation, for the dependent resources, they will have to be created separately.

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
EffectiveRoutesContent content = new EffectiveRoutesContent
{
    ResourceId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/expressRouteGateways/expressRouteGatewayName/expressRouteConnections/connectionName"),
    VirtualWanResourceType = "ExpressRouteConnection",
};
ArmOperation<VirtualHubEffectiveRouteList> lro = await virtualHub.GetVirtualHubEffectiveRoutesAsync(WaitUntil.Completed, content: content);
VirtualHubEffectiveRouteList result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
