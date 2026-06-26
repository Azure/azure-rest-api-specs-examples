const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets a list of database restore points.
 *
 * @summary gets a list of database restore points.
 * x-ms-original-file: 2025-01-01/DataWarehouseRestorePointsListByDatabase.json
 */
async function listDatawarehouseDatabaseRestorePoints() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.restorePoints.listByDatabase(
    "Default-SQL-SouthEastAsia",
    "testserver",
    "testDatabase",
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
