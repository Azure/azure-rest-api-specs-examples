const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a sync member.
 *
 * @summary Creates or updates a sync member.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/SyncMemberCreate.json
 */
async function createANewSyncMember() {
  const subscriptionId =
    process.env["SQL_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SQL_RESOURCE_GROUP"] || "syncgroupcrud-65440";
  const serverName = "syncgroupcrud-8475";
  const databaseName = "syncgroupcrud-4328";
  const syncGroupName = "syncgroupcrud-3187";
  const syncMemberName = "syncmembercrud-4879";
  const parameters = {
    databaseName: "syncgroupcrud-7421",
    databaseType: "AzureSqlDatabase",
    serverName: "syncgroupcrud-3379.database.windows.net",
    syncDirection: "Bidirectional",
    syncMemberAzureDatabaseResourceId:
      "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/syncgroupcrud-65440/providers/Microsoft.Sql/servers/syncgroupcrud-8475/databases/syncgroupcrud-4328",
    usePrivateLinkConnection: true,
    userName: "myUser",
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.syncMembers.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serverName,
    databaseName,
    syncGroupName,
    syncMemberName,
    parameters,
  );
  console.log(result);
}
