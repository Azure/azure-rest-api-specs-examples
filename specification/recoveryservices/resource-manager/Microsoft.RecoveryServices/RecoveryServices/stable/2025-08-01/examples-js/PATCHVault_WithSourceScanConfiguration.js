const { RecoveryServicesClient } = require("@azure/arm-recoveryservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates the vault.
 *
 * @summary updates the vault.
 * x-ms-original-file: 2025-08-01/PATCHVault_WithSourceScanConfiguration.json
 */
async function updateVaultWithSourceScanConfiguration() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "77777777-b0c6-47a2-b37c-d8e65a629c18";
  const client = new RecoveryServicesClient(credential, subscriptionId);
  const result = await client.vaults.update("HelloWorld", "swaggerExample", {
    identity: {
      type: "UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/85bf5e8c-3084-4f42-add2-746ebb7e97b2/resourcegroups/defaultrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/examplemsi":
          {},
      },
    },
    properties: {
      securitySettings: {
        sourceScanConfiguration: {
          sourceScanIdentity: {
            operationIdentityType: "UserAssigned",
            userAssignedIdentity:
              "/subscriptions/85bf5e8c-3084-4f42-add2-746ebb7e97b2/resourcegroups/defaultrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/examplemsi",
          },
          state: "Enabled",
        },
      },
    },
    tags: { PatchKey: "PatchKeyUpdated" },
  });
  console.log(result);
}
