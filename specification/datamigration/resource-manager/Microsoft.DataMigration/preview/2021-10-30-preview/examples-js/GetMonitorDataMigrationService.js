const { DataMigrationManagementClient } = require("@azure/arm-datamigration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve the Monitoring Data.
 *
 * @summary Retrieve the Monitoring Data.
 * x-ms-original-file: specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2021-10-30-preview/examples/GetMonitorDataMigrationService.json
 */
async function retrieveTheMonitoringData() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "testrg";
  const sqlMigrationServiceName = "service1";
  const credential = new DefaultAzureCredential();
  const client = new DataMigrationManagementClient(credential, subscriptionId);
  const result = await client.sqlMigrationServices.listMonitoringData(
    resourceGroupName,
    sqlMigrationServiceName
  );
  console.log(result);
}

retrieveTheMonitoringData().catch(console.error);
