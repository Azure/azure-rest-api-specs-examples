const { DataMigrationManagementClient } = require("@azure/arm-datamigration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to retrieve the specified database migration for a given SQL VM.
 *
 * @summary retrieve the specified database migration for a given SQL VM.
 * x-ms-original-file: 2025-09-01-preview/SqlVmGetDatabaseMigrationExpanded.json
 */
async function getSqlVMDatabaseMigrationWithTheExpandParameter() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new DataMigrationManagementClient(credential, subscriptionId);
  const result = await client.databaseMigrationsSqlVm.get("testrg", "testvm", "db1", {
    expand: "MigrationStatusDetails",
  });
  console.log(result);
}
