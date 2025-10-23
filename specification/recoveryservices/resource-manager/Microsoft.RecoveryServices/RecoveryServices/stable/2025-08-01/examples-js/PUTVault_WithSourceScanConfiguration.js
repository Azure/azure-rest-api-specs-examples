const { RecoveryServicesClient } = require("@azure/arm-recoveryservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a Recovery Services vault.
 *
 * @summary creates or updates a Recovery Services vault.
 * x-ms-original-file: 2025-08-01/PUTVault_WithSourceScanConfiguration.json
 */
async function createOrUpdateVaultWithSourceScanConfiguration() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "77777777-b0c6-47a2-b37c-d8e65a629c18";
  const client = new RecoveryServicesClient(credential, subscriptionId);
  const result = await client.vaults.createOrUpdate(
    "Default-RecoveryServices-ResourceGroup",
    "swaggerExample",
    {
      identity: { type: "SystemAssigned" },
      location: "West US",
      properties: {
        publicNetworkAccess: "Enabled",
        securitySettings: {
          sourceScanConfiguration: {
            sourceScanIdentity: { operationIdentityType: "SystemAssigned" },
            state: "Enabled",
          },
        },
      },
      sku: { name: "Standard" },
    },
  );
  console.log(result);
}
