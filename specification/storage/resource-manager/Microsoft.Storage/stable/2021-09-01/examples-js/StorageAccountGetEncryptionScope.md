Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storage_17.2.0/sdk/storage/arm-storage/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns the properties for the specified encryption scope.
 *
 * @summary Returns the properties for the specified encryption scope.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/StorageAccountGetEncryptionScope.json
 */
async function storageAccountGetEncryptionScope() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "resource-group-name";
  const accountName = "{storage-account-name}";
  const encryptionScopeName = "{encryption-scope-name}";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.encryptionScopes.get(
    resourceGroupName,
    accountName,
    encryptionScopeName
  );
  console.log(result);
}

storageAccountGetEncryptionScope().catch(console.error);
```
