using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/P2sVpnGatewaysDisconnectP2sVpnConnections.json
// this example is just showing the usage of "P2sVpnGateways_DisconnectP2SVpnConnections" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this P2SVpnGatewayResource created on azure
// for more information of creating P2SVpnGatewayResource, please refer to the document of P2SVpnGatewayResource
string subscriptionId = "subid";
string resourceGroupName = "p2s-vpn-gateway-test";
string p2sVpnGatewayName = "p2svpngateway";
ResourceIdentifier p2sVpnGatewayResourceId = P2SVpnGatewayResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, p2sVpnGatewayName);
P2SVpnGatewayResource p2sVpnGateway = client.GetP2SVpnGatewayResource(p2sVpnGatewayResourceId);

// invoke the operation
P2SVpnConnectionRequest request = new P2SVpnConnectionRequest
{
    VpnConnectionIds = { "vpnconnId1", "vpnconnId2" },
};
await p2sVpnGateway.DisconnectP2SVpnConnectionsAsync(WaitUntil.Completed, request);

Console.WriteLine("Succeeded");
