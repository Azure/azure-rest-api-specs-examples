const { PostgreSQLManagementFlexibleServerClient } = require("@azure/arm-postgresql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Gets details of a migration.
 *
 * @summary Gets details of a migration.
 * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/Migrations_GetMigrationWithSuccessfulValidationButMigrationFailure.json
 */
async function migrationsGetMigrationWithSuccessfulValidationButMigrationFailure() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = process.env["POSTGRESQL_RESOURCE_GROUP"] || "testrg";
  const targetDbServerName = "testtarget";
  const migrationName = "testmigrationwithsuccessfulvalidationbutmigrationfailure";
  const credential = new DefaultAzureCredential();
  const client = new PostgreSQLManagementFlexibleServerClient(credential);
  const result = await client.migrations.get(
    subscriptionId,
    resourceGroupName,
    targetDbServerName,
    migrationName,
  );
  console.log(result);
}
