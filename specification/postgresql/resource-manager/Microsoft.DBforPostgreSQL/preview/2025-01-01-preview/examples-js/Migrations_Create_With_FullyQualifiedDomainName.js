const { PostgreSQLManagementFlexibleServerClient } = require("@azure/arm-postgresql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates a new migration.
 *
 * @summary Creates a new migration.
 * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/Migrations_Create_With_FullyQualifiedDomainName.json
 */
async function migrationsCreateWithFullyQualifiedDomainName() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = process.env["POSTGRESQL_RESOURCE_GROUP"] || "testrg";
  const targetDbServerName = "testtarget";
  const migrationName = "testmigration";
  const parameters = {
    dbsToMigrate: ["db1", "db2", "db3", "db4"],
    location: "westus",
    migrationMode: "Offline",
    overwriteDbsInTarget: "True",
    secretParameters: {
      adminCredentials: {
        sourceServerPassword: "xxxxxxxx",
        targetServerPassword: "xxxxxxxx",
      },
    },
    sourceDbServerFullyQualifiedDomainName: "testsourcefqdn.example.com",
    sourceDbServerResourceId:
      "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBForPostgreSql/servers/testsource",
    targetDbServerFullyQualifiedDomainName: "test-target-fqdn.example.com",
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
