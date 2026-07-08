const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets the specified virtual network peering.
 *
 * @summary gets the specified virtual network peering.
 * x-ms-original-file: 2025-07-01/VirtualNetworkV6SubnetPeeringGet.json
 */
async function getV6SubnetPeering() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworkPeerings.get("peerTest", "vnet1", "peer");
  console.log(result);
}
