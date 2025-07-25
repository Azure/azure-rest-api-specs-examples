const { DataMigrationManagementClient } = require("@azure/arm-datamigration");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Stop on going migration for the database.
 *
 * @summary Stop on going migration for the database.
 * x-ms-original-file: specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2025-03-15-preview/examples/SqlDbCancelDatabaseMigration.json
 */
async function stopOngoingMigrationForTheDatabase() {
  const subscriptionId =
    process.env["DATAMIGRATION_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["DATAMIGRATION_RESOURCE_GROUP"] || "testrg";
  const sqlDbInstanceName = "sqldbinstance";
  const targetDbName = "db1";
  const parameters = {
    migrationOperationId: "9a90bb84-e70f-46f7-b0ae-1aef5b3b9f07",
  };
  const credential = new DefaultAzureCredential();
  const client = new DataMigrationManagementClient(credential, subscriptionId);
  const result = await client.databaseMigrationsSqlDb.beginCancelAndWait(
    resourceGroupName,
    sqlDbInstanceName,
    targetDbName,
    parameters,
  );
  console.log(result);
}
