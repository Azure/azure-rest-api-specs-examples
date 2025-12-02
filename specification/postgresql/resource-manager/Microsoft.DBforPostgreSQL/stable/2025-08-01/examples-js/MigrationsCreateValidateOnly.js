const { PostgreSQLManagementFlexibleServerClient } = require("@azure/arm-postgresql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates a new migration.
 *
 * @summary Creates a new migration.
 * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/MigrationsCreateValidateOnly.json
 */
async function createAMigrationForValidatingOnly() {
  const subscriptionId =
    process.env["POSTGRESQL_SUBSCRIPTION_ID"] || "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = process.env["POSTGRESQL_RESOURCE_GROUP"] || "exampleresourcegroup";
  const serverName = "exampleserver";
  const migrationName = "examplemigration";
  const parameters = {
    dbsToMigrate: ["exampledatabase1", "exampledatabase2", "exampledatabase3", "exampledatabase4"],
    location: "eastus",
    migrationMode: "Offline",
    migrationOption: "Validate",
    overwriteDbsInTarget: "True",
    secretParameters: {
      adminCredentials: {
        sourceServerPassword: "examplesourcepassword",
        targetServerPassword: "exampletargetpassword",
      },
    },
    sourceDbServerResourceId:
      "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.DBForPostgreSql/servers/examplesource",
  };
  const credential = new DefaultAzureCredential();
  const client = new PostgreSQLManagementFlexibleServerClient(credential, subscriptionId);
  const result = await client.migrations.create(
    resourceGroupName,
    serverName,
    migrationName,
    parameters,
  );
  console.log(result);
}
