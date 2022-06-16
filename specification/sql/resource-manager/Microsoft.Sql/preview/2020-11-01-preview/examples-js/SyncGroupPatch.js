const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a sync group.
 *
 * @summary Updates a sync group.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/SyncGroupPatch.json
 */
async function updateASyncGroup() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "syncgroupcrud-65440";
  const serverName = "syncgroupcrud-8475";
  const databaseName = "syncgroupcrud-4328";
  const syncGroupName = "syncgroupcrud-3187";
  const parameters = {
    conflictResolutionPolicy: "HubWin",
    hubDatabasePassword: "hubPassword",
    hubDatabaseUserName: "hubUser",
    interval: -1,
    syncDatabaseId:
      "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/syncgroupcrud-3521/providers/Microsoft.Sql/servers/syncgroupcrud-8475/databases/syncgroupcrud-4328",
    usePrivateLinkConnection: true,
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.syncGroups.beginUpdateAndWait(
    resourceGroupName,
    serverName,
    databaseName,
    syncGroupName,
    parameters
  );
  console.log(result);
}

updateASyncGroup().catch(console.error);
