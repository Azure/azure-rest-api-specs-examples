const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new database or updates an existing database.
 *
 * @summary creates a new database or updates an existing database.
 * x-ms-original-file: 2025-01-01/ManagedDatabaseCreatePointInTimeRestore.json
 */
async function createsANewManagedDatabaseUsingPointInTimeRestore() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.managedDatabases.createOrUpdate(
    "Default-SQL-SouthEastAsia",
    "managedInstance",
    "managedDatabase",
    {
      location: "southeastasia",
      createMode: "PointInTimeRestore",
      restorePointInTime: new Date("2017-07-14T05:35:31.503Z"),
      sourceDatabaseId:
        "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/managedInstances/testsvr/databases/testdb",
    },
  );
  console.log(result);
}
