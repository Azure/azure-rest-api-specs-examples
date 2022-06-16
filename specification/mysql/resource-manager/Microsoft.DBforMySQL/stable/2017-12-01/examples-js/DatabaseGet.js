const { MySQLManagementClient } = require("@azure/arm-mysql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets information about a database.
 *
 * @summary Gets information about a database.
 * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/DatabaseGet.json
 */
async function databaseGet() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "TestGroup";
  const serverName = "testserver";
  const databaseName = "db1";
  const credential = new DefaultAzureCredential();
  const client = new MySQLManagementClient(credential, subscriptionId);
  const result = await client.databases.get(resourceGroupName, serverName, databaseName);
  console.log(result);
}

databaseGet().catch(console.error);
