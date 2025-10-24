const { RecoveryServicesClient } = require("@azure/arm-recoveryservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates the vault.
 *
 * @summary updates the vault.
 * x-ms-original-file: 2025-08-01/PATCHVault_WithRedundancySettings.json
 */
async function updateVaultWithRedundancySetting() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "77777777-b0c6-47a2-b37c-d8e65a629c18";
  const client = new RecoveryServicesClient(credential, subscriptionId);
  const result = await client.vaults.update("HelloWorld", "swaggerExample", {
    properties: {
      redundancySettings: {
        crossRegionRestore: "Enabled",
        standardTierStorageRedundancy: "GeoRedundant",
      },
    },
  });
  console.log(result);
}
