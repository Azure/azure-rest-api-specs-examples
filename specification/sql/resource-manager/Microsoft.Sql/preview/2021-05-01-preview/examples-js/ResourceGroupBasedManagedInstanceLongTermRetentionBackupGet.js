const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a long term retention backup for a managed database.
 *
 * @summary Gets a long term retention backup for a managed database.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/ResourceGroupBasedManagedInstanceLongTermRetentionBackupGet.json
 */
async function getTheLongTermRetentionBackup() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "testResourceGroup";
  const locationName = "japaneast";
  const managedInstanceName = "testInstance";
  const databaseName = "testDatabase";
  const backupName = "55555555-6666-7777-8888-999999999999;131637960820000000";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.longTermRetentionManagedInstanceBackups.getByResourceGroup(
    resourceGroupName,
    locationName,
    managedInstanceName,
    databaseName,
    backupName
  );
  console.log(result);
}

getTheLongTermRetentionBackup().catch(console.error);
