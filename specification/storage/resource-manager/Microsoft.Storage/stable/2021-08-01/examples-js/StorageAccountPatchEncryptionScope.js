const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function storageAccountPatchEncryptionScope() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "resource-group-name";
  const accountName = "{storage-account-name}";
  const encryptionScopeName = "{encryption-scope-name}";
  const encryptionScope = {
    keyVaultProperties: {
      keyUri: "https://testvault.vault.core.windows.net/keys/key1/863425f1358359c",
    },
    source: "Microsoft.KeyVault",
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.encryptionScopes.patch(
    resourceGroupName,
    accountName,
    encryptionScopeName,
    encryptionScope
  );
  console.log(result);
}

storageAccountPatchEncryptionScope().catch(console.error);
