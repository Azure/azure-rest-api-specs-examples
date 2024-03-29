const { DataMigrationManagementClient } = require("@azure/arm-datamigration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Stop migrations in progress for the database
 *
 * @summary Stop migrations in progress for the database
 * x-ms-original-file: specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2021-10-30-preview/examples/SqlMiCancelDatabaseMigration.json
 */
async function stopOngoingMigrationForTheDatabase() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "testrg";
  const managedInstanceName = "managedInstance1";
  const targetDbName = "db1";
  const parameters = {
    migrationOperationId: "4124fe90-d1b6-4b50-b4d9-46d02381f59a",
  };
  const credential = new DefaultAzureCredential();
  const client = new DataMigrationManagementClient(credential, subscriptionId);
  const result = await client.databaseMigrationsSqlMi.beginCancelAndWait(
    resourceGroupName,
    managedInstanceName,
    targetDbName,
    parameters
  );
  console.log(result);
}

stopOngoingMigrationForTheDatabase().catch(console.error);
