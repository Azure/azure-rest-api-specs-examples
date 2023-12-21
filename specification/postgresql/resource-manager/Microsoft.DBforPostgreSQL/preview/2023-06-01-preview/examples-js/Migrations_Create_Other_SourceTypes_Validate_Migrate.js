const { PostgreSQLManagementFlexibleServerClient } = require("@azure/arm-postgresql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new migration.
 *
 * @summary Creates a new migration.
 * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-06-01-preview/examples/Migrations_Create_Other_SourceTypes_Validate_Migrate.json
 */
async function createMigrationWithOtherSourceTypesForValidateAndMigrate() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = process.env["POSTGRESQL_RESOURCE_GROUP"] || "testrg";
  const targetDbServerName = "testtarget";
  const migrationName = "testmigration";
  const parameters = {
    dbsToMigrate: ["db1", "db2", "db3", "db4"],
    location: "westus",
    migrationMode: "Offline",
    migrationOption: "ValidateAndMigrate",
    overwriteDbsInTarget: "True",
    secretParameters: {
      adminCredentials: {
        sourceServerPassword: "xxxxxxxx",
        targetServerPassword: "xxxxxxxx",
      },
    },
    sourceDbServerResourceId: "testsource:5432@pguser",
    sourceType: "OnPremises",
    sslMode: "Prefer",
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
