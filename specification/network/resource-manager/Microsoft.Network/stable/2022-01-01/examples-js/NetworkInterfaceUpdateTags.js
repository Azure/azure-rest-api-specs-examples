const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a network interface tags.
 *
 * @summary Updates a network interface tags.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkInterfaceUpdateTags.json
 */
async function updateNetworkInterfaceTags() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const networkInterfaceName = "test-nic";
  const parameters = { tags: { tag1: "value1", tag2: "value2" } };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.networkInterfaces.updateTags(
    resourceGroupName,
    networkInterfaceName,
    parameters
  );
  console.log(result);
}

updateNetworkInterfaceTags().catch(console.error);
