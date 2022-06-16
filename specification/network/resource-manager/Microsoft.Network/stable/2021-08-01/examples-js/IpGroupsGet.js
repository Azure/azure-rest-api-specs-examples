const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified ipGroups.
 *
 * @summary Gets the specified ipGroups.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/IpGroupsGet.json
 */
async function getIPGroups() {
  const subscriptionId = "subId";
  const resourceGroupName = "myResourceGroup";
  const ipGroupsName = "ipGroups1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.ipGroups.get(resourceGroupName, ipGroupsName);
  console.log(result);
}

getIPGroups().catch(console.error);
