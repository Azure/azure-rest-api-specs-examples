const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a peering in the specified virtual network.
 *
 * @summary Creates or updates a peering in the specified virtual network.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/VirtualNetworkV6SubnetPeeringCreate.json
 */
async function createV6SubnetPeering() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "peerTest";
  const virtualNetworkName = "vnet1";
  const virtualNetworkPeeringName = "peer";
  const virtualNetworkPeeringParameters = {
    allowForwardedTraffic: true,
    allowGatewayTransit: false,
    allowVirtualNetworkAccess: true,
    enableOnlyIPv6Peering: true,
    localSubnetNames: ["Subnet1", "Subnet4"],
    peerCompleteVnets: false,
    remoteSubnetNames: ["Subnet2"],
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
    virtualNetworkPeeringParameters,
  );
  console.log(result);
}
