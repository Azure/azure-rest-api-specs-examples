const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to list managed database columns
 *
 * @summary list managed database columns
 * x-ms-original-file: 2025-01-01/ManagedColumnsListByDatabaseMax.json
 */
async function filterManagedDatabaseColumns() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.managedDatabaseColumns.listByDatabase(
    "myRG",
    "serverName",
    "myDatabase",
    {
      schema: ["dbo"],
      table: ["customer", "address"],
      column: ["username"],
      orderBy: ["schema asc", "table", "column desc"],
    },
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
