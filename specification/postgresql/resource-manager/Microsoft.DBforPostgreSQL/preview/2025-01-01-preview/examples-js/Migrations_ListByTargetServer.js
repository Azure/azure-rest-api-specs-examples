const { PostgreSQLManagementFlexibleServerClient } = require("@azure/arm-postgresql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to List all the migrations on a given target server.
 *
 * @summary List all the migrations on a given target server.
 * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/Migrations_ListByTargetServer.json
 */
async function migrationsListByTargetServer() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = process.env["POSTGRESQL_RESOURCE_GROUP"] || "testrg";
  const targetDbServerName = "testtarget";
  const migrationListFilter = "All";
  const options = {
    migrationListFilter,
  };
  const credential = new DefaultAzureCredential();
  const client = new PostgreSQLManagementFlexibleServerClient(credential);
  const resArray = new Array();
  for await (const item of client.migrations.listByTargetServer(
    subscriptionId,
    resourceGroupName,
    targetDbServerName,
    options,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
