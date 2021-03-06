const { PostgreSQLManagementClient } = require("@azure/arm-postgresql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a PostgreSQL Server key.
 *
 * @summary Gets a PostgreSQL Server key.
 * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2020-01-01/examples/ServerKeyGet.json
 */
async function getThePostgreSqlServerKey() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "testrg";
  const serverName = "testserver";
  const keyName = "someVault_someKey_01234567890123456789012345678901";
  const credential = new DefaultAzureCredential();
  const client = new PostgreSQLManagementClient(credential, subscriptionId);
  const result = await client.serverKeys.get(resourceGroupName, serverName, keyName);
  console.log(result);
}

getThePostgreSqlServerKey().catch(console.error);
