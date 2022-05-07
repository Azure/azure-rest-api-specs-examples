Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storage_17.2.0/sdk/storage/arm-storage/README.md) on how to add the SDK to your project and authenticate.

```javascript
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
```
