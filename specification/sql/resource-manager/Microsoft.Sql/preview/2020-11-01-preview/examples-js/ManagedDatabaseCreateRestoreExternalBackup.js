const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new database or updates an existing database.
 *
 * @summary Creates a new database or updates an existing database.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedDatabaseCreateRestoreExternalBackup.json
 */
async function createsANewManagedDatabaseByRestoringFromAnExternalBackup() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "Default-SQL-SouthEastAsia";
  const managedInstanceName = "managedInstance";
  const databaseName = "managedDatabase";
  const parameters = {
    autoCompleteRestore: true,
    collation: "SQL_Latin1_General_CP1_CI_AS",
    createMode: "RestoreExternalBackup",
    lastBackupName: "last_backup_name",
    location: "southeastasia",
    storageContainerSasToken: "sv=2015-12-11&sr=c&sp=rl&sig=1234",
    storageContainerUri: "https://myaccountname.blob.core.windows.net/backups",
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.managedDatabases.beginCreateOrUpdateAndWait(
    resourceGroupName,
    managedInstanceName,
    databaseName,
    parameters
  );
  console.log(result);
}

createsANewManagedDatabaseByRestoringFromAnExternalBackup().catch(console.error);
