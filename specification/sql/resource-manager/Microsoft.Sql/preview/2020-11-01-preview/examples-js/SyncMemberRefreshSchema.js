const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Refreshes a sync member database schema.
 *
 * @summary Refreshes a sync member database schema.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/SyncMemberRefreshSchema.json
 */
async function refreshASyncMemberDatabaseSchema() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "syncgroupcrud-65440";
  const serverName = "syncgroupcrud-8475";
  const databaseName = "syncgroupcrud-4328";
  const syncGroupName = "syncgroupcrud-3187";
  const syncMemberName = "syncgroupcrud-4879";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.syncMembers.beginRefreshMemberSchemaAndWait(
    resourceGroupName,
    serverName,
    databaseName,
    syncGroupName,
    syncMemberName
  );
  console.log(result);
}

refreshASyncMemberDatabaseSchema().catch(console.error);
