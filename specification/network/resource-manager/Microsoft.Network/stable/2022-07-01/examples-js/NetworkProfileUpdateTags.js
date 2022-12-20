const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates network profile tags.
 *
 * @summary Updates network profile tags.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/NetworkProfileUpdateTags.json
 */
async function updateNetworkProfileTags() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const networkProfileName = "test-np";
  const parameters = { tags: { tag1: "value1", tag2: "value2" } };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.networkProfiles.updateTags(
    resourceGroupName,
    networkProfileName,
    parameters
  );
  console.log(result);
}

updateNetworkProfileTags().catch(console.error);
