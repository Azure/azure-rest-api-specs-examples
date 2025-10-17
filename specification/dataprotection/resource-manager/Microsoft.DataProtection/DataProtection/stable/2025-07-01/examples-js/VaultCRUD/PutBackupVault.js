const { DataProtectionClient } = require("@azure/arm-dataprotection");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a BackupVault resource belonging to a resource group.
 *
 * @summary creates or updates a BackupVault resource belonging to a resource group.
 * x-ms-original-file: 2025-07-01/VaultCRUD/PutBackupVault.json
 */
async function createBackupVault() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "0b352192-dcac-4cc7-992e-a96190ccc68c";
  const client = new DataProtectionClient(credential, subscriptionId);
  const result = await client.backupVaults.createOrUpdate("SampleResourceGroup", "swaggerExample", {
    location: "WestUS",
    properties: {
      featureSettings: { crossRegionRestoreSettings: { state: "Enabled" } },
      monitoringSettings: {
        azureMonitorAlertSettings: { alertsForAllJobFailures: "Enabled" },
      },
      securitySettings: {
        softDeleteSettings: { retentionDurationInDays: 14, state: "Enabled" },
      },
      storageSettings: [{ type: "LocallyRedundant", datastoreType: "VaultStore" }],
    },
    tags: { key1: "val1" },
  });
  console.log(result);
}
