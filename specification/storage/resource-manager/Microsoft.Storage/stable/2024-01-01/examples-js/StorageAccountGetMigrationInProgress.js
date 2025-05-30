const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Gets the status of the ongoing migration for the specified storage account.
 *
 * @summary Gets the status of the ongoing migration for the specified storage account.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/StorageAccountGetMigrationInProgress.json
 */
async function storageAccountGetMigrationInProgress() {
  const subscriptionId = process.env["STORAGE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["STORAGE_RESOURCE_GROUP"] || "resource-group-name";
  const accountName = "accountname";
  const migrationName = "default";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.storageAccounts.getCustomerInitiatedMigration(
    resourceGroupName,
    accountName,
    migrationName,
  );
  console.log(result);
}
