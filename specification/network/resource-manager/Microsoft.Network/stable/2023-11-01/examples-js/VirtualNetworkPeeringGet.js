const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified virtual network peering.
 *
 * @summary Gets the specified virtual network peering.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-11-01/examples/VirtualNetworkPeeringGet.json
 */
async function getPeering() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "peerTest";
  const virtualNetworkName = "vnet1";
  const virtualNetworkPeeringName = "peer";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworkPeerings.get(
    resourceGroupName,
    virtualNetworkName,
    virtualNetworkPeeringName,
  );
  console.log(result);
}
