const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Synchronously creates or updates an encryption scope under the specified storage account. If an encryption scope is already created and a subsequent request is issued with different properties, the encryption scope properties will be updated per the specified request.
 *
 * @summary Synchronously creates or updates an encryption scope under the specified storage account. If an encryption scope is already created and a subsequent request is issued with different properties, the encryption scope properties will be updated per the specified request.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2022-05-01/examples/StorageAccountPutEncryptionScopeWithInfrastructureEncryption.json
 */
async function storageAccountPutEncryptionScopeWithInfrastructureEncryption() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "resource-group-name";
  const accountName = "{storage-account-name}";
  const encryptionScopeName = "{encryption-scope-name}";
  const encryptionScope = {
    requireInfrastructureEncryption: true,
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.encryptionScopes.put(
    resourceGroupName,
    accountName,
    encryptionScopeName,
    encryptionScope
  );
  console.log(result);
}

storageAccountPutEncryptionScopeWithInfrastructureEncryption().catch(console.error);
