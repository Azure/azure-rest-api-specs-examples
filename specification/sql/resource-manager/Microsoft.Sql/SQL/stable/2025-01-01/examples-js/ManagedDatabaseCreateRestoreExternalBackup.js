const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new database or updates an existing database.
 *
 * @summary creates a new database or updates an existing database.
 * x-ms-original-file: 2025-01-01/ManagedDatabaseCreateRestoreExternalBackup.json
 */
async function createsANewManagedDatabaseByRestoringFromAnExternalBackup() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.managedDatabases.createOrUpdate(
    "Default-SQL-SouthEastAsia",
    "managedInstance",
    "managedDatabase",
    {
      location: "southeastasia",
      autoCompleteRestore: true,
      collation: "SQL_Latin1_General_CP1_CI_AS",
      createMode: "RestoreExternalBackup",
      lastBackupName: "last_backup_name",
      storageContainerSasToken: "sv=2015-12-11&sr=c&sp=rl&sig=1234",
      storageContainerUri: "https://myaccountname.blob.core.windows.net/backups",
    },
  );
  console.log(result);
}
