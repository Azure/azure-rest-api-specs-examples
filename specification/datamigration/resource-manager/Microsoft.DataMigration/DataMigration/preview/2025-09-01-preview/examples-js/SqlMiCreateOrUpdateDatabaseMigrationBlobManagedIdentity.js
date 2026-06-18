const { DataMigrationManagementClient } = require("@azure/arm-datamigration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a new database migration to a given SQL Managed Instance.
 *
 * @summary create a new database migration to a given SQL Managed Instance.
 * x-ms-original-file: 2025-09-01-preview/SqlMiCreateOrUpdateDatabaseMigrationBlobManagedIdentity.json
 */
async function createOrUpdateDatabaseMigrationResourceFromAzureBlobUsingManagedIdentity() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new DataMigrationManagementClient(credential, subscriptionId);
  const result = await client.databaseMigrationsSqlMi.createOrUpdate(
    "testrg",
    "managedInstance1",
    "db1",
    {
      properties: {
        backupConfiguration: {
          sourceLocation: {
            azureBlob: {
              authType: "ManagedIdentity",
              blobContainerName: "test",
              identity: {
                type: "UserAssigned",
                userAssignedIdentities: {
                  "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testidentity":
                    {},
                },
              },
              storageAccountResourceId:
                "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Storage/storageAccounts/teststorageaccount",
            },
          },
        },
        kind: "SqlMi",
        migrationService:
          "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.DataMigration/sqlMigrationServices/testagent",
        offlineConfiguration: { lastBackupName: "last_backup_file_name", offline: true },
        scope:
          "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/instance",
        sourceDatabaseName: "aaa",
        sqlServerInstanceId:
          "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.AzureArcData/SqlServerInstances/instanceName",
      },
    },
  );
  console.log(result);
}
