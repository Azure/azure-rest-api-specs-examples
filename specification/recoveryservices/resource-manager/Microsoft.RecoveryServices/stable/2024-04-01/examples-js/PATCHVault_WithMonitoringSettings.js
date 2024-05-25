const { RecoveryServicesClient } = require("@azure/arm-recoveryservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the vault.
 *
 * @summary Updates the vault.
 * x-ms-original-file: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2024-04-01/examples/PATCHVault_WithMonitoringSettings.json
 */
async function updateVaultWithMonitoringSetting() {
  const subscriptionId =
    process.env["RECOVERYSERVICES_SUBSCRIPTION_ID"] || "77777777-b0c6-47a2-b37c-d8e65a629c18";
  const resourceGroupName = process.env["RECOVERYSERVICES_RESOURCE_GROUP"] || "HelloWorld";
  const vaultName = "swaggerExample";
  const vault = {
    properties: {
      monitoringSettings: {
        azureMonitorAlertSettings: {
          alertsForAllFailoverIssues: "Disabled",
          alertsForAllJobFailures: "Enabled",
          alertsForAllReplicationIssues: "Enabled",
        },
        classicAlertSettings: {
          alertsForCriticalOperations: "Disabled",
          emailNotificationsForSiteRecovery: "Enabled",
        },
      },
    },
    tags: { patchKey: "PatchKeyUpdated" },
  };
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesClient(credential, subscriptionId);
  const result = await client.vaults.beginUpdateAndWait(resourceGroupName, vaultName, vault);
  console.log(result);
}
