const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Completes the restore operation on a managed database.
 *
 * @summary Completes the restore operation on a managed database.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedDatabaseCompleteExternalRestore.json
 */
async function completesAManagedDatabaseExternalBackupRestore() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "myRG";
  const managedInstanceName = "myManagedInstanceName";
  const databaseName = "myDatabase";
  const parameters = {
    lastBackupName: "testdb1_log4",
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.managedDatabases.beginCompleteRestoreAndWait(
    resourceGroupName,
    managedInstanceName,
    databaseName,
    parameters
  );
  console.log(result);
}

completesAManagedDatabaseExternalBackupRestore().catch(console.error);
