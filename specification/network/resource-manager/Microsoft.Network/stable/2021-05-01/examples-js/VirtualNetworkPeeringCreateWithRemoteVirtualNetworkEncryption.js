const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a peering in the specified virtual network.
 *
 * @summary Creates or updates a peering in the specified virtual network.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VirtualNetworkPeeringCreateWithRemoteVirtualNetworkEncryption.json
 */
async function createPeeringWithRemoteVirtualNetworkEncryption() {
  const subscriptionId = "subid";
  const resourceGroupName = "peerTest";
  const virtualNetworkName = "vnet1";
  const virtualNetworkPeeringName = "peer";
  const virtualNetworkPeeringParameters = {
    allowForwardedTraffic: true,
    allowGatewayTransit: false,
    allowVirtualNetworkAccess: true,
    remoteVirtualNetwork: {
      id: "/subscriptions/subid/resourceGroups/peerTest/providers/Microsoft.Network/virtualNetworks/vnet2",
    },
    useRemoteGateways: false,
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworkPeerings.beginCreateOrUpdateAndWait(
    resourceGroupName,
    virtualNetworkName,
    virtualNetworkPeeringName,
    virtualNetworkPeeringParameters
  );
  console.log(result);
}

createPeeringWithRemoteVirtualNetworkEncryption().catch(console.error);
