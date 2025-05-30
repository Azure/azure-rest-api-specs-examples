const { PostgreSQLManagementFlexibleServerClient } = require("@azure/arm-postgresql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates a new migration.
 *
 * @summary Creates a new migration.
 * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2024-08-01/examples/Migrations_Create_With_Other_Users.json
 */
async function migrationsCreateByPassingUserNames() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = process.env["POSTGRESQL_RESOURCE_GROUP"] || "testrg";
  const targetDbServerName = "testtarget";
  const migrationName = "testmigration";
  const parameters = {
    dbsToMigrate: ["db1", "db2", "db3", "db4"],
    location: "westus",
    migrationMode: "Offline",
    secretParameters: {
      adminCredentials: {
        sourceServerPassword: "xxxxxxxx",
        targetServerPassword: "xxxxxxxx",
      },
      sourceServerUsername: "newadmin@testsource",
      targetServerUsername: "targetadmin",
    },
    sourceDbServerResourceId:
      "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBForPostgreSql/servers/testsource",
  };
  const credential = new DefaultAzureCredential();
  const client = new PostgreSQLManagementFlexibleServerClient(credential);
  const result = await client.migrations.create(
    subscriptionId,
    resourceGroupName,
    targetDbServerName,
    migrationName,
    parameters,
  );
  console.log(result);
}
