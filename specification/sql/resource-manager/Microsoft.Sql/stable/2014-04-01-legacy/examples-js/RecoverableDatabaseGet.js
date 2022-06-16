const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a recoverable database, which is a resource representing a database's geo backup
 *
 * @summary Gets a recoverable database, which is a resource representing a database's geo backup
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01-legacy/examples/RecoverableDatabaseGet.json
 */
async function getARecoverableDatabase() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "recoverabledatabasetest-6852";
  const serverName = "recoverabledatabasetest-2080";
  const databaseName = "recoverabledatabasetest-9187";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.recoverableDatabases.get(resourceGroupName, serverName, databaseName);
  console.log(result);
}

getARecoverableDatabase().catch(console.error);
