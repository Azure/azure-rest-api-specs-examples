const { RecoveryServicesClient } = require("@azure/arm-recoveryservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates the vault.
 *
 * @summary updates the vault.
 * x-ms-original-file: 2025-08-01/PatchVault_WithCMK2.json
 */
async function updateResourceWithCustomerManagedKeys2() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "77777777-b0c6-47a2-b37c-d8e65a629c18";
  const client = new RecoveryServicesClient(credential, subscriptionId);
  const result = await client.vaults.update("HelloWorld", "swaggerExample", {
    identity: { type: "SystemAssigned" },
    properties: {
      encryption: { kekIdentity: { useSystemAssignedIdentity: true } },
    },
    tags: { PatchKey: "PatchKeyUpdated" },
  });
  console.log(result);
}
