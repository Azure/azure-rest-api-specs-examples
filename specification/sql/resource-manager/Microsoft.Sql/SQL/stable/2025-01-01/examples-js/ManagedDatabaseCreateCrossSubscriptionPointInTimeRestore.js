const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new database or updates an existing database.
 *
 * @summary creates a new database or updates an existing database.
 * x-ms-original-file: 2025-01-01/ManagedDatabaseCreateCrossSubscriptionPointInTimeRestore.json
 */
async function createsANewManagedDatabaseUsingCrossSubscriptionPointInTimeRestore() {
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
      crossSubscriptionSourceDatabaseId:
        "/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/managedInstances/testsvr2/databases/testdb",
      crossSubscriptionTargetManagedInstanceId:
        "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/managedInstances/testsvr",
      restorePointInTime: new Date("2017-07-14T05:35:31.503Z"),
    },
  );
  console.log(result);
}
