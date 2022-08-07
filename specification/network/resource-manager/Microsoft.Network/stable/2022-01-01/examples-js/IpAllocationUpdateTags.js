const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a IpAllocation tags.
 *
 * @summary Updates a IpAllocation tags.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/IpAllocationUpdateTags.json
 */
async function updateVirtualNetworkTags() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const ipAllocationName = "test-ipallocation";
  const parameters = { tags: { tag1: "value1", tag2: "value2" } };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.ipAllocations.updateTags(
    resourceGroupName,
    ipAllocationName,
    parameters
  );
  console.log(result);
}

updateVirtualNetworkTags().catch(console.error);
