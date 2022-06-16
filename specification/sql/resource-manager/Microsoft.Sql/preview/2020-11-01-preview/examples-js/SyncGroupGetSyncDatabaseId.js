const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a collection of sync database ids.
 *
 * @summary Gets a collection of sync database ids.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/SyncGroupGetSyncDatabaseId.json
 */
async function getASyncDatabaseId() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const locationName = "westus";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.syncGroups.listSyncDatabaseIds(locationName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getASyncDatabaseId().catch(console.error);
