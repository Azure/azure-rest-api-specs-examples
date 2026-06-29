const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the update operation can be used to update the SKU, encryption, access tier, or tags for a storage account. It can also be used to map the account to a custom domain. Only one custom domain is supported per storage account; the replacement/change of custom domain is not supported. In order to replace an old custom domain, the old value must be cleared/unregistered before a new value can be set. The update of multiple properties is supported. This call does not change the storage keys for the account. If you want to change the storage account keys, use the regenerate keys operation. The location and name of the storage account cannot be changed after creation.
 *
 * @summary the update operation can be used to update the SKU, encryption, access tier, or tags for a storage account. It can also be used to map the account to a custom domain. Only one custom domain is supported per storage account; the replacement/change of custom domain is not supported. In order to replace an old custom domain, the old value must be cleared/unregistered before a new value can be set. The update of multiple properties is supported. This call does not change the storage keys for the account. If you want to change the storage account keys, use the regenerate keys operation. The location and name of the storage account cannot be changed after creation.
 * x-ms-original-file: 2026-04-01/StorageAccountUpdateUserAssignedIdentityWithFederatedIdentityClientId.json
 */
async function storageAccountUpdateUserAssignedIdentityWithFederatedIdentityClientId() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.storageAccounts.update("res131918", "sto131918", {
    identity: {
      type: "UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/{subscription-id}/resourceGroups/res9101/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{managed-identity-name}":
          {},
      },
    },
    kind: "Storage",
    encryption: {
      encryptionIdentity: {
        encryptionFederatedIdentityClientId: "3109d1c4-a5de-4d84-8832-feabb916a4b6",
        encryptionUserAssignedIdentity:
          "/subscriptions/{subscription-id}/resourceGroups/res9101/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{managed-identity-name}",
      },
      keySource: "Microsoft.Keyvault",
      keyVaultProperties: {
        keyName: "wrappingKey",
        keyVaultUri: "https://myvault8569.vault.azure.net",
        keyVersion: "",
      },
      services: {
        blob: { enabled: true, keyType: "Account" },
        file: { enabled: true, keyType: "Account" },
      },
    },
    sku: { name: "Standard_LRS" },
  });
  console.log(result);
}
