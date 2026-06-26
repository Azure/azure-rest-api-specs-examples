const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new database or updates an existing database.
 *
 * @summary creates a new database or updates an existing database.
 * x-ms-original-file: 2025-01-01/CreateDwDatabaseCrossSubscriptionRestore.json
 */
async function createsADataWarehouseDatabaseAsACrossSubscriptionRestoreFromABackupOfADroppedDatabase() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.databases.createOrUpdate(
    "Default-SQL-SouthEastAsia",
    "testsvr",
    "testdw",
    {
      location: "southeastasia",
      createMode: "Restore",
      sourceResourceId:
        "/subscriptions/55555555-6666-7777-8888-999999999999/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/srcsvr/restorableDroppedDatabases/srcdw,131403269876900000",
    },
  );
  console.log(result);
}
