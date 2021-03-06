const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function storageAccountUpdateUserAssignedIdentityWithFederatedIdentityClientId() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res131918";
  const accountName = "sto131918";
  const parameters = {
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
    identity: {
      type: "UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/{subscriptionId}/resourceGroups/res9101/providers/MicrosoftManagedIdentity/userAssignedIdentities/{managedIdentityName}":
          {},
      },
    },
    kind: "Storage",
    sku: { name: "Standard_LRS" },
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.storageAccounts.update(resourceGroupName, accountName, parameters);
  console.log(result);
}

storageAccountUpdateUserAssignedIdentityWithFederatedIdentityClientId().catch(console.error);
