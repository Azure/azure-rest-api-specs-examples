const { RecoveryServicesClient } = require("@azure/arm-recoveryservices");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Updates the vault.
 *
 * @summary Updates the vault.
 * x-ms-original-file: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/PATCHVault_WithSourceScanConfiguration.json
 */
async function updateVaultWithSourceScanConfiguration() {
  const subscriptionId =
    process.env["RECOVERYSERVICES_SUBSCRIPTION_ID"] || "77777777-b0c6-47a2-b37c-d8e65a629c18";
  const resourceGroupName = process.env["RECOVERYSERVICES_RESOURCE_GROUP"] || "HelloWorld";
  const vaultName = "swaggerExample";
  const vault = {
    identity: {
      type: "UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/85bf5e8c30844f42Add2746ebb7e97b2/resourcegroups/defaultrg/providers/MicrosoftManagedIdentity/userAssignedIdentities/examplemsi":
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
    tags: { patchKey: "PatchKeyUpdated" },
  };
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesClient(credential, subscriptionId);
  const result = await client.vaults.beginUpdateAndWait(resourceGroupName, vaultName, vault);
  console.log(result);
}
