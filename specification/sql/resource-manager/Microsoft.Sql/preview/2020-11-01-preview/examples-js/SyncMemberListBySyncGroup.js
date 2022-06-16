const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists sync members in the given sync group.
 *
 * @summary Lists sync members in the given sync group.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/SyncMemberListBySyncGroup.json
 */
async function listSyncMembersUnderASyncGroup() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "syncgroupcrud-65440";
  const serverName = "syncgroupcrud-8475";
  const databaseName = "syncgroupcrud-4328";
  const syncGroupName = "syncgroupcrud-3187";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.syncMembers.listBySyncGroup(
    resourceGroupName,
    serverName,
    databaseName,
    syncGroupName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listSyncMembersUnderASyncGroup().catch(console.error);
