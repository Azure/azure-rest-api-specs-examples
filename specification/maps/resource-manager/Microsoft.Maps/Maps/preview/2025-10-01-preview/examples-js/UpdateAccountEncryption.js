const { AzureMapsManagementClient } = require("@azure/arm-maps");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates a Maps Account. Only a subset of the parameters may be updated after creation, such as Sku, Tags, Properties.
 *
 * @summary updates a Maps Account. Only a subset of the parameters may be updated after creation, such as Sku, Tags, Properties.
 * x-ms-original-file: 2025-10-01-preview/UpdateAccountEncryption.json
 */
async function updateAccountEncryption() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "21a9967a-e8a9-4656-a70b-96ff1c4d05a0";
  const client = new AzureMapsManagementClient(credential, subscriptionId);
  const result = await client.accounts.update("myResourceGroup", "myMapsAccount", {
    identity: {
      type: "SystemAssigned",
      userAssignedIdentities: {
        "/subscriptions/21a9967a-e8a9-4656-a70b-96ff1c4d05a0/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identityName":
          {},
      },
    },
    encryption: {
      customerManagedKeyEncryption: {
        keyEncryptionKeyIdentity: { identityType: "systemAssignedIdentity" },
        keyEncryptionKeyUrl: "https://contosovault.vault.azure.net/keys/contosokek",
      },
    },
  });
  console.log(result);
}
