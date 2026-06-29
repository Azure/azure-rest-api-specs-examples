const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to asynchronously creates a new storage account with the specified parameters. If an account is already created and a subsequent create request is issued with different properties, the account properties will be updated. If an account is already created and a subsequent create or update request is issued with the exact same set of properties, the request will succeed.
 *
 * @summary asynchronously creates a new storage account with the specified parameters. If an account is already created and a subsequent create request is issued with different properties, the account properties will be updated. If an account is already created and a subsequent create or update request is issued with the exact same set of properties, the request will succeed.
 * x-ms-original-file: 2026-04-01/StorageAccountCreateUserAssignedEncryptionIdentityWithCMK.json
 */
async function storageAccountCreateUserAssignedEncryptionIdentityWithCMK() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.storageAccounts.create("res9101", "sto4445", {
    identity: {
      type: "UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/{subscription-id}/resourceGroups/res9101/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{managed-identity-name}":
          {},
      },
    },
    kind: "Storage",
    location: "eastus",
    encryption: {
      encryptionIdentity: {
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
