const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Gets the specified virtual network peering.
 *
 * @summary Gets the specified virtual network peering.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/VirtualNetworkPeeringGetWithRemoteVirtualNetworkEncryption.json
 */
async function getPeeringWithRemoteVirtualNetworkEncryption() {
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
