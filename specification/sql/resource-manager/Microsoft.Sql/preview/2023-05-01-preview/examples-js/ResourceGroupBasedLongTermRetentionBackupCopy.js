const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Copy an existing long term retention backup to a different server.
 *
 * @summary Copy an existing long term retention backup to a different server.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2023-05-01-preview/examples/ResourceGroupBasedLongTermRetentionBackupCopy.json
 */
async function copyTheLongTermRetentionBackup() {
  const subscriptionId =
    process.env["SQL_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SQL_RESOURCE_GROUP"] || "testResourceGroup";
  const locationName = "japaneast";
  const longTermRetentionServerName = "testserver";
  const longTermRetentionDatabaseName = "testDatabase";
  const backupName = "55555555-6666-7777-8888-999999999999;131637960820000000";
  const parameters = {
    targetBackupStorageRedundancy: "Geo",
    targetDatabaseName: "testDatabase2",
    targetServerResourceId:
      "/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/resourceGroups/resourceGroup/servers/testserver2",
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.longTermRetentionBackups.beginCopyByResourceGroupAndWait(
    resourceGroupName,
    locationName,
    longTermRetentionServerName,
    longTermRetentionDatabaseName,
    backupName,
    parameters,
  );
  console.log(result);
}
