const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List managed database columns
 *
 * @summary List managed database columns
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedDatabaseColumnListByTable.json
 */
async function listManagedDatabaseColumns() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "myRG";
  const managedInstanceName = "myManagedInstanceName";
  const databaseName = "myDatabase";
  const schemaName = "dbo";
  const tableName = "table1";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.managedDatabaseColumns.listByTable(
    resourceGroupName,
    managedInstanceName,
    databaseName,
    schemaName,
    tableName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listManagedDatabaseColumns().catch(console.error);
