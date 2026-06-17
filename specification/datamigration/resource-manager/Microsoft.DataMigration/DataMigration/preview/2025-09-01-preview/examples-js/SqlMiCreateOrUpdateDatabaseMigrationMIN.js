const { DataMigrationManagementClient } = require("@azure/arm-datamigration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a new database migration to a given SQL Managed Instance.
 *
 * @summary create a new database migration to a given SQL Managed Instance.
 * x-ms-original-file: 2025-09-01-preview/SqlMiCreateOrUpdateDatabaseMigrationMIN.json
 */
async function createOrUpdateDatabaseMigrationResourceWithMinimumParameters() {
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
            fileShare: { path: "C:\\aaa\\bbb\\ccc", password: "placeholder", username: "name" },
          },
          targetLocation: {
            accountKey: "abcd",
            storageAccountResourceId: "account.database.windows.net",
          },
        },
        kind: "SqlMi",
        migrationService:
          "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.DataMigration/sqlMigrationServices/testagent",
        scope:
          "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/instance",
        sourceDatabaseName: "aaa",
        sourceSqlConnection: {
          authentication: "WindowsAuthentication",
          dataSource: "aaa",
          encryptConnection: true,
          password: "placeholder",
          trustServerCertificate: true,
          userName: "bbb",
        },
        sqlServerInstanceId:
          "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.AzureArcData/SqlServerInstances/instanceName",
      },
    },
  );
  console.log(result);
}
