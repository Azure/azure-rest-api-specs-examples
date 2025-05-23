const { PostgreSQLManagementFlexibleServerClient } = require("@azure/arm-postgresql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to This method checks whether a proposed migration name is valid and available.
 *
 * @summary This method checks whether a proposed migration name is valid and available.
 * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/CheckMigrationNameAvailability.json
 */
async function checkMigrationNameAvailability() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = process.env["POSTGRESQL_RESOURCE_GROUP"] || "testrg";
  const targetDbServerName = "testtarget";
  const parameters = {
    name: "name1",
    type: "Microsoft.DBforPostgreSQL/flexibleServers/migrations",
  };
  const credential = new DefaultAzureCredential();
  const client = new PostgreSQLManagementFlexibleServerClient(credential);
  const result = await client.checkMigrationNameAvailability(
    subscriptionId,
    resourceGroupName,
    targetDbServerName,
    parameters,
  );
  console.log(result);
}
