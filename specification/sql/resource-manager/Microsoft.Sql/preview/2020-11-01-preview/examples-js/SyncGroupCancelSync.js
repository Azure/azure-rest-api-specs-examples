const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Cancels a sync group synchronization.
 *
 * @summary Cancels a sync group synchronization.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/SyncGroupCancelSync.json
 */
async function cancelASyncGroupSynchronization() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "syncgroupcrud-65440";
  const serverName = "syncgroupcrud-8475";
  const databaseName = "syncgroupcrud-4328";
  const syncGroupName = "syncgroupcrud-3187";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.syncGroups.cancelSync(
    resourceGroupName,
    serverName,
    databaseName,
    syncGroupName
  );
  console.log(result);
}

cancelASyncGroupSynchronization().catch(console.error);
