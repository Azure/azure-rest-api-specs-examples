const { NetworkManagementClient } = require("@azure/arm-network-profile-2020-09-01-hybrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The GetBgpPeerStatus operation retrieves the status of all BGP peers.
 *
 * @summary The GetBgpPeerStatus operation retrieves the status of all BGP peers.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2018-11-01/examples/VirtualNetworkGatewayGetBGPPeerStatus.json
 */
async function getVirtualNetworkGatewayBgpPeerStatus() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const virtualNetworkGatewayName = "vpngw";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworkGateways.beginGetBgpPeerStatusAndWait(
    resourceGroupName,
    virtualNetworkGatewayName
  );
  console.log(result);
}
