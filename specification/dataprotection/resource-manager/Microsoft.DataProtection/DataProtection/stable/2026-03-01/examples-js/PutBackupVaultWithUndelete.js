const { DataProtectionClient } = require("@azure/arm-dataprotection");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a BackupVault resource belonging to a resource group.
 *
 * @summary creates or updates a BackupVault resource belonging to a resource group.
 * x-ms-original-file: 2026-03-01/PutBackupVaultWithUndelete.json
 */
async function restoreASoftDeletedBackupVault() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new DataProtectionClient(credential, subscriptionId);
  const result = await client.backupVaults.createOrUpdate(
    "SampleResourceGroup",
    "swaggerExample",
    {
      location: "WestUS",
      properties: {
        monitoringSettings: { azureMonitorAlertSettings: { alertsForAllJobFailures: "Enabled" } },
        securitySettings: {
          softDeleteSettings: { retentionDurationInDays: 14, state: "Enabled" },
          immutabilitySettings: { state: "Disabled" },
        },
        storageSettings: [{ datastoreType: "VaultStore", type: "LocallyRedundant" }],
        featureSettings: {
          crossSubscriptionRestoreSettings: { state: "Disabled" },
          crossRegionRestoreSettings: { state: "Enabled" },
        },
      },
      tags: { key1: "val1" },
    },
    {
      xMsDeletedVaultId:
        "/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.DataProtection/locations/WestUS/deletedVaults/swaggerExample",
    },
  );
  console.log(result);
}
