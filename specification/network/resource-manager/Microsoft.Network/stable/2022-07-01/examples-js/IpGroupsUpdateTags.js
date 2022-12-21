const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates tags of an IpGroups resource.
 *
 * @summary Updates tags of an IpGroups resource.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/IpGroupsUpdateTags.json
 */
async function updateIPGroups() {
  const subscriptionId = "subId";
  const resourceGroupName = "myResourceGroup";
  const ipGroupsName = "ipGroups1";
  const parameters = { tags: { key1: "value1", key2: "value2" } };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.ipGroups.updateGroups(resourceGroupName, ipGroupsName, parameters);
  console.log(result);
}

updateIPGroups().catch(console.error);
