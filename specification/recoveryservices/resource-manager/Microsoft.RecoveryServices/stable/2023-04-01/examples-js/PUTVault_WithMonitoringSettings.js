const { RecoveryServicesClient } = require("@azure/arm-recoveryservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a Recovery Services vault.
 *
 * @summary Creates or updates a Recovery Services vault.
 * x-ms-original-file: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2023-04-01/examples/PUTVault_WithMonitoringSettings.json
 */
async function createOrUpdateVaultWithMonitoringSetting() {
  const subscriptionId =
    process.env["RECOVERYSERVICES_SUBSCRIPTION_ID"] || "77777777-b0c6-47a2-b37c-d8e65a629c18";
  const resourceGroupName =
    process.env["RECOVERYSERVICES_RESOURCE_GROUP"] || "Default-RecoveryServices-ResourceGroup";
  const vaultName = "swaggerExample";
  const vault = {
    identity: { type: "SystemAssigned" },
    location: "West US",
    properties: {
      monitoringSettings: {
        azureMonitorAlertSettings: { alertsForAllJobFailures: "Enabled" },
        classicAlertSettings: { alertsForCriticalOperations: "Disabled" },
      },
      publicNetworkAccess: "Enabled",
    },
    sku: { name: "Standard" },
  };
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesClient(credential, subscriptionId);
  const result = await client.vaults.beginCreateOrUpdateAndWait(
    resourceGroupName,
    vaultName,
    vault
  );
  console.log(result);
}
