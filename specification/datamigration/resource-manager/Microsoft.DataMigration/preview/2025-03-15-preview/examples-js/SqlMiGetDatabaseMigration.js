const { DataMigrationManagementClient } = require("@azure/arm-datamigration");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Retrieve the specified database migration for a given SQL Managed Instance.
 *
 * @summary Retrieve the specified database migration for a given SQL Managed Instance.
 * x-ms-original-file: specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2025-03-15-preview/examples/SqlMiGetDatabaseMigration.json
 */
async function getSqlMiDatabaseMigrationWithoutTheExpandParameter() {
  const subscriptionId =
    process.env["DATAMIGRATION_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["DATAMIGRATION_RESOURCE_GROUP"] || "testrg";
  const managedInstanceName = "managedInstance1";
  const targetDbName = "db1";
  const credential = new DefaultAzureCredential();
  const client = new DataMigrationManagementClient(credential, subscriptionId);
  const result = await client.databaseMigrationsSqlMi.get(
    resourceGroupName,
    managedInstanceName,
    targetDbName,
  );
  console.log(result);
}
