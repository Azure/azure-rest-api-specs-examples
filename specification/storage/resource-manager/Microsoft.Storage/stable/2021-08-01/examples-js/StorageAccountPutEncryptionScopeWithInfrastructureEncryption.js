const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

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
