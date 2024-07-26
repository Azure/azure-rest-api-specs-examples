const { DataProtectionClient } = require("@azure/arm-dataprotection");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a BackupVault resource belonging to a resource group. For example, updating tags for a resource.
 *
 * @summary Updates a BackupVault resource belonging to a resource group. For example, updating tags for a resource.
 * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2024-04-01/examples/VaultCRUD/PatchBackupVaultWithCMK.json
 */
async function patchBackupVaultWithCmk() {
  const subscriptionId =
    process.env["DATAPROTECTION_SUBSCRIPTION_ID"] || "0b352192-dcac-4cc7-992e-a96190ccc68c";
  const resourceGroupName = process.env["DATAPROTECTION_RESOURCE_GROUP"] || "SampleResourceGroup";
  const vaultName = "swaggerExample";
  const parameters = {
    properties: {
      monitoringSettings: {
        azureMonitorAlertSettings: { alertsForAllJobFailures: "Enabled" },
      },
      securitySettings: {
        encryptionSettings: {
          infrastructureEncryption: "Enabled",
          kekIdentity: { identityType: "SystemAssigned" },
          keyVaultProperties: {
            keyUri: "https://cmk2xkv.vault.azure.net/keys/Key1/0767b348bb1a4c07baa6c4ec0055d2b3",
          },
          state: "Enabled",
        },
        immutabilitySettings: { state: "Disabled" },
        softDeleteSettings: { retentionDurationInDays: 90, state: "On" },
      },
    },
    tags: { newKey: "newVal" },
  };
  const credential = new DefaultAzureCredential();
  const client = new DataProtectionClient(credential, subscriptionId);
  const result = await client.backupVaults.beginUpdateAndWait(
    resourceGroupName,
    vaultName,
    parameters,
  );
  console.log(result);
}
