const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of operations performed on the database.
 *
 * @summary Gets a list of operations performed on the database.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2021-02-01-preview/examples/ListDatabaseOperations.json
 */
async function listTheDatabaseManagementOperations() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "sqlcrudtest-7398";
  const serverName = "sqlcrudtest-4645";
  const databaseName = "testdb";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.databaseOperations.listByDatabase(
    resourceGroupName,
    serverName,
    databaseName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listTheDatabaseManagementOperations().catch(console.error);
