const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified ipGroups.
 *
 * @summary Deletes the specified ipGroups.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/IpGroupsDelete.json
 */
async function deleteIPGroups() {
  const subscriptionId = "subId";
  const resourceGroupName = "myResourceGroup";
  const ipGroupsName = "ipGroups1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.ipGroups.beginDeleteAndWait(resourceGroupName, ipGroupsName);
  console.log(result);
}

deleteIPGroups().catch(console.error);
