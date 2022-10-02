const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an VirtualNetworkTap tags.
 *
 * @summary Updates an VirtualNetworkTap tags.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/VirtualNetworkTapUpdateTags.json
 */
async function updateVirtualNetworkTapTags() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const tapName = "test-vtap";
  const tapParameters = {
    tags: { tag1: "value1", tag2: "value2" },
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworkTaps.updateTags(
    resourceGroupName,
    tapName,
    tapParameters
  );
  console.log(result);
}

updateVirtualNetworkTapTags().catch(console.error);
