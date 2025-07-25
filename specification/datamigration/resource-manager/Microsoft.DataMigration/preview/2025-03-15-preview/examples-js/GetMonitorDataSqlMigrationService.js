const { DataMigrationManagementClient } = require("@azure/arm-datamigration");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Retrieve the registered Integration Runtime nodes and their monitoring data for a given Database Migration Service.
 *
 * @summary Retrieve the registered Integration Runtime nodes and their monitoring data for a given Database Migration Service.
 * x-ms-original-file: specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2025-03-15-preview/examples/GetMonitorDataSqlMigrationService.json
 */
async function retrieveTheMonitoringData() {
  const subscriptionId =
    process.env["DATAMIGRATION_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["DATAMIGRATION_RESOURCE_GROUP"] || "testrg";
  const sqlMigrationServiceName = "service1";
  const credential = new DefaultAzureCredential();
  const client = new DataMigrationManagementClient(credential, subscriptionId);
  const result = await client.sqlMigrationServices.listMonitoringData(
    resourceGroupName,
    sqlMigrationServiceName,
  );
  console.log(result);
}
