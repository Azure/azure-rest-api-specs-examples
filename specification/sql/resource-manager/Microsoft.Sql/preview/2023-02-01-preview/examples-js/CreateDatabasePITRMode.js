const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates a new database or updates an existing database.
 *
 * @summary Creates a new database or updates an existing database.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2023-02-01-preview/examples/CreateDatabasePITRMode.json
 */
async function createsADatabaseFromPointInTimeRestore() {
  const subscriptionId =
    process.env["SQL_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SQL_RESOURCE_GROUP"] || "Default-SQL-SouthEastAsia";
  const serverName = "testsvr";
  const databaseName = "dbpitr";
  const parameters = {
    createMode: "PointInTimeRestore",
    location: "southeastasia",
    restorePointInTime: new Date("2020-10-22T05:35:31.503Z"),
    sourceDatabaseId:
      "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SoutheastAsia/providers/Microsoft.Sql/servers/testsvr/databases/testdb",
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.databases.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serverName,
    databaseName,
    parameters,
  );
  console.log(result);
}
