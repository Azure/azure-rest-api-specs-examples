const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets database usages.
 *
 * @summary Gets database usages.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2021-02-01-preview/examples/GetDatabaseUsages.json
 */
async function getsDatabaseUsages() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "Default-SQL-SouthEastAsia";
  const serverName = "testsvr";
  const databaseName = "testdb";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.databaseUsages.listByDatabase(
    resourceGroupName,
    serverName,
    databaseName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

getsDatabaseUsages().catch(console.error);
