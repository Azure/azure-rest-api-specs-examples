const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified virtual network peering.
 *
 * @summary Deletes the specified virtual network peering.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/VirtualNetworkPeeringDelete.json
 */
async function deletePeering() {
  const subscriptionId = "subid";
  const resourceGroupName = "peerTest";
  const virtualNetworkName = "vnet1";
  const virtualNetworkPeeringName = "peer";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworkPeerings.beginDeleteAndWait(
    resourceGroupName,
    virtualNetworkName,
    virtualNetworkPeeringName
  );
  console.log(result);
}

deletePeering().catch(console.error);
