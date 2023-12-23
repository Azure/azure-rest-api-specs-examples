const { DataProtectionClient } = require("@azure/arm-dataprotection");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a BackupVault resource belonging to a resource group.
 *
 * @summary Creates or updates a BackupVault resource belonging to a resource group.
 * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2023-11-01/examples/VaultCRUD/PutBackupVaultWithMSI.json
 */
async function createBackupVaultWithMsi() {
  const subscriptionId =
    process.env["DATAPROTECTION_SUBSCRIPTION_ID"] || "0b352192-dcac-4cc7-992e-a96190ccc68c";
  const resourceGroupName = process.env["DATAPROTECTION_RESOURCE_GROUP"] || "SampleResourceGroup";
  const vaultName = "swaggerExample";
  const parameters = {
    identity: { type: "systemAssigned" },
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
  };
  const credential = new DefaultAzureCredential();
  const client = new DataProtectionClient(credential, subscriptionId);
  const result = await client.backupVaults.beginCreateOrUpdateAndWait(
    resourceGroupName,
    vaultName,
    parameters,
  );
  console.log(result);
}
