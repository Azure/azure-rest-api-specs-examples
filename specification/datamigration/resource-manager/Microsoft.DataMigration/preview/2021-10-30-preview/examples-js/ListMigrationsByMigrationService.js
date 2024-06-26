const { DataMigrationManagementClient } = require("@azure/arm-datamigration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve the List of database migrations attached to the service.
 *
 * @summary Retrieve the List of database migrations attached to the service.
 * x-ms-original-file: specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2021-10-30-preview/examples/ListMigrationsByMigrationService.json
 */
async function listDatabaseMigrationsAttachedToTheService() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "testrg";
  const sqlMigrationServiceName = "service1";
  const credential = new DefaultAzureCredential();
  const client = new DataMigrationManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.sqlMigrationServices.listMigrations(
    resourceGroupName,
    sqlMigrationServiceName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listDatabaseMigrationsAttachedToTheService().catch(console.error);
