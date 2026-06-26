const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to list database columns
 *
 * @summary list database columns
 * x-ms-original-file: 2025-01-01/ColumnsListByDatabaseMin.json
 */
async function listDatabaseColumns() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.databaseColumns.listByDatabase(
    "myRG",
    "serverName",
    "myDatabase",
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
