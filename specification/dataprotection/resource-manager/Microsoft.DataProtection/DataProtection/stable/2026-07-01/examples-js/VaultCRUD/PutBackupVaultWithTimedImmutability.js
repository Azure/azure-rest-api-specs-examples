const { DataProtectionClient } = require("@azure/arm-dataprotection");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a BackupVault resource belonging to a resource group.
 *
 * @summary creates or updates a BackupVault resource belonging to a resource group.
 * x-ms-original-file: 2026-07-01/VaultCRUD/PutBackupVaultWithTimedImmutability.json
 */
async function createBackupVaultWithTimedImmutability() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "04cf684a-d41f-4550-9f70-7708a3a2283b";
  const client = new DataProtectionClient(credential, subscriptionId);
  const result = await client.backupVaults.createOrUpdate("SampleResourceGroup", "swaggerExample", {
    location: "WestUS",
    properties: {
      securitySettings: {
        immutabilitySettings: {
          state: "Unlocked",
          configuration: { type: "TimeBased", durationInDays: 30 },
        },
        softDeleteSettings: { retentionDurationInDays: 14, state: "On" },
      },
      storageSettings: [{ type: "LocallyRedundant", datastoreType: "VaultStore" }],
    },
    tags: { key1: "val1" },
    identity: { type: "None" },
  });
  console.log(result);
}
