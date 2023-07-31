using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Network;
using Azure.ResourceManager.Network.Models;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/examples/VirtualNetworkGatewaysDisconnectP2sVpnConnections.json
// this example is just showing the usage of "VirtualNetworkGateways_DisconnectVirtualNetworkGatewayVpnConnections" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualNetworkGatewayResource created on azure
// for more information of creating VirtualNetworkGatewayResource, please refer to the document of VirtualNetworkGatewayResource
string subscriptionId = "subid";
string resourceGroupName = "vpn-gateway-test";
string virtualNetworkGatewayName = "vpngateway";
ResourceIdentifier virtualNetworkGatewayResourceId = VirtualNetworkGatewayResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualNetworkGatewayName);
VirtualNetworkGatewayResource virtualNetworkGateway = client.GetVirtualNetworkGatewayResource(virtualNetworkGatewayResourceId);

// invoke the operation
P2SVpnConnectionRequest request = new P2SVpnConnectionRequest()
{
    VpnConnectionIds =
    {
    "vpnconnId1","vpnconnId2"
    },
};
await virtualNetworkGateway.DisconnectVirtualNetworkGatewayVpnConnectionsAsync(WaitUntil.Completed, request);

Console.WriteLine($"Succeeded");
