const { DataMigrationManagementClient } = require("@azure/arm-datamigration");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Get Database Migration resources for the scope.
 *
 * @summary Get Database Migration resources for the scope.
 * x-ms-original-file: specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2025-03-15-preview/examples/CosmosDbMongoListByScopeDatabaseMigration.json
 */
async function getMongoToCosmosDbMongoVCoreDatabaseMigrationWithoutTheExpandParameter() {
  const subscriptionId =
    process.env["DATAMIGRATION_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["DATAMIGRATION_RESOURCE_GROUP"] || "testrg";
  const targetResourceName = "targetCosmosDbClusterName";
  const credential = new DefaultAzureCredential();
  const client = new DataMigrationManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.databaseMigrationsMongoToCosmosDbvCoreMongo.listForScope(
    resourceGroupName,
    targetResourceName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
