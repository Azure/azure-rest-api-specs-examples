const { DataProtectionClient } = require("@azure/arm-dataprotection");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a BackupVault resource belonging to a resource group.
 *
 * @summary creates or updates a BackupVault resource belonging to a resource group.
 * x-ms-original-file: 2026-06-01/VaultCRUD/PutBackupVaultCMKSettings_ResourceGuardEnabled.json
 */
async function createOrUpdateBackupVaultWithCMKAndResourceGuardEnabled() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "0b352192-dcac-4cc7-992e-a96190ccc68c";
  const client = new DataProtectionClient(credential, subscriptionId);
  const result = await client.backupVaults.createOrUpdate("SampleResourceGroup", "swaggerExample", {
    location: "WestUS",
    tags: { key1: "val1" },
    identity: { type: "None" },
    properties: {
      monitoringSettings: { azureMonitorAlertSettings: { alertsForAllJobFailures: "Enabled" } },
      securitySettings: {
        softDeleteSettings: { state: "Off", retentionDurationInDays: 0 },
        immutabilitySettings: { state: "Disabled" },
        encryptionSettings: {
          state: "Enabled",
          keyVaultProperties: {
            keyUri: "https://cmk2xkv.vault.azure.net/keys/Key1/0767b348bb1a4c07baa6c4ec0055d2b3",
          },
          kekIdentity: {
            identityType: "UserAssigned",
            identityId:
              "/subscriptions/85bf5e8c-3084-4f42-add2-746ebb7e97b2/resourcegroups/defaultrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/examplemsi",
          },
          infrastructureEncryption: "Enabled",
        },
      },
      storageSettings: [{ datastoreType: "VaultStore", type: "LocallyRedundant" }],
    },
  });
  console.log(result);
}
