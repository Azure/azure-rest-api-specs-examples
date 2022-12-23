const { StorageManagementClient } = require("@azure/arm-storage-profile-2020-09-01-hybrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update encryption scope properties as specified in the request body. Update fails if the specified encryption scope does not already exist.
 *
 * @summary Update encryption scope properties as specified in the request body. Update fails if the specified encryption scope does not already exist.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2019-06-01/examples/StorageAccountPatchEncryptionScope.json
 */
async function storageAccountPatchEncryptionScope() {
  const subscriptionId = process.env["STORAGE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["STORAGE_RESOURCE_GROUP"] || "resource-group-name";
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
