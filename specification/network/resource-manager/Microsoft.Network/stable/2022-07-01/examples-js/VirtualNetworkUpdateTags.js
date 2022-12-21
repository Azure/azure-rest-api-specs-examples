const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a virtual network tags.
 *
 * @summary Updates a virtual network tags.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/VirtualNetworkUpdateTags.json
 */
async function updateVirtualNetworkTags() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const virtualNetworkName = "test-vnet";
  const parameters = { tags: { tag1: "value1", tag2: "value2" } };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworks.updateTags(
    resourceGroupName,
    virtualNetworkName,
    parameters
  );
  console.log(result);
}

updateVirtualNetworkTags().catch(console.error);
