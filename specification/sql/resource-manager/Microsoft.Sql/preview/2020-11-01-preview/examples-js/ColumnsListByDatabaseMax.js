const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List database columns
 *
 * @summary List database columns
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ColumnsListByDatabaseMax.json
 */
async function filterDatabaseColumns() {
  const subscriptionId =
    process.env["SQL_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SQL_RESOURCE_GROUP"] || "myRG";
  const serverName = "serverName";
  const databaseName = "myDatabase";
  const schema = ["dbo"];
  const table = ["customer", "address"];
  const column = ["username"];
  const orderBy = ["schema asc", "table", "column desc"];
  const options = {
    schema,
    table,
    column,
    orderBy,
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.databaseColumns.listByDatabase(
    resourceGroupName,
    serverName,
    databaseName,
    options,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
