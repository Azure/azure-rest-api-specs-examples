const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to synchronously creates or updates an encryption scope under the specified storage account. If an encryption scope is already created and a subsequent request is issued with different properties, the encryption scope properties will be updated per the specified request.
 *
 * @summary synchronously creates or updates an encryption scope under the specified storage account. If an encryption scope is already created and a subsequent request is issued with different properties, the encryption scope properties will be updated per the specified request.
 * x-ms-original-file: 2026-04-01/StorageAccountPutEncryptionScope.json
 */
async function storageAccountPutEncryptionScope() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.encryptionScopes.put(
    "resource-group-name",
    "accountname",
    "{encryption-scope-name}",
    {},
  );
  console.log(result);
}
