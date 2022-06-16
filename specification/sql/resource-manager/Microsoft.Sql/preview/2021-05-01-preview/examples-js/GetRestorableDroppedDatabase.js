const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a restorable dropped database.
 *
 * @summary Gets a restorable dropped database.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/GetRestorableDroppedDatabase.json
 */
async function getsARestorableDroppedDatabase() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "Default-SQL-SouthEastAsia";
  const serverName = "testsvr";
  const restorableDroppedDatabaseId = "testdb,131403269876900000";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.restorableDroppedDatabases.get(
    resourceGroupName,
    serverName,
    restorableDroppedDatabaseId
  );
  console.log(result);
}

getsARestorableDroppedDatabase().catch(console.error);
