const { DataMigrationManagementClient } = require("@azure/arm-datamigration");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Retrieve the Database Migration resource.
 *
 * @summary Retrieve the Database Migration resource.
 * x-ms-original-file: specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2025-03-15-preview/examples/SqlDbGetDatabaseMigrationExpanded.json
 */
async function getSqlDbDatabaseMigrationWithTheExpandParameter() {
  const subscriptionId =
    process.env["DATAMIGRATION_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["DATAMIGRATION_RESOURCE_GROUP"] || "testrg";
  const sqlDbInstanceName = "sqldbinstance";
  const targetDbName = "db1";
  const expand = "MigrationStatusDetails";
  const options = { expand };
  const credential = new DefaultAzureCredential();
  const client = new DataMigrationManagementClient(credential, subscriptionId);
  const result = await client.databaseMigrationsSqlDb.get(
    resourceGroupName,
    sqlDbInstanceName,
    targetDbName,
    options,
  );
  console.log(result);
}
