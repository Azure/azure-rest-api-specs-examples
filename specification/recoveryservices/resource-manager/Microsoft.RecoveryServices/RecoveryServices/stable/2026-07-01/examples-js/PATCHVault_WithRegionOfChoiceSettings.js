const { RecoveryServicesClient } = require("@azure/arm-recoveryservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates the vault.
 *
 * @summary updates the vault.
 * x-ms-original-file: 2026-07-01/PATCHVault_WithRegionOfChoiceSettings.json
 */
async function updateVaultWithRegionOfChoiceSettings() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "77777777-b0c6-47a2-b37c-d8e65a629c18";
  const client = new RecoveryServicesClient(credential, subscriptionId);
  const result = await client.vaults.update("HelloWorld", "swaggerExample", {
    properties: { regionOfChoiceSettings: { status: "Enabled" } },
    tags: { PatchKey: "PatchKeyUpdated" },
  });
  console.log(result);
}
