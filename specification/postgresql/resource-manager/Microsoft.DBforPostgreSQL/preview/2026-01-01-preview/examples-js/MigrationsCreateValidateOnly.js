const { PostgreSQLManagementFlexibleServerClient } = require("@azure/arm-postgresql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new migration.
 *
 * @summary creates a new migration.
 * x-ms-original-file: 2026-01-01-preview/MigrationsCreateValidateOnly.json
 */
async function createAMigrationForValidatingOnly() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const client = new PostgreSQLManagementFlexibleServerClient(credential, subscriptionId);
  const result = await client.migrations.create(
    "exampleresourcegroup",
    "exampleserver",
    "examplemigration",
    {
      location: "eastus",
      dbsToMigrate: [
        "exampledatabase1",
        "exampledatabase2",
        "exampledatabase3",
        "exampledatabase4",
      ],
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
    },
  );
  console.log(result);
}
