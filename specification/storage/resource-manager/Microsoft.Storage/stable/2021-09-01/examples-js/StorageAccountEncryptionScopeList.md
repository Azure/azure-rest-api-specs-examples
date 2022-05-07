Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storage_17.2.0/sdk/storage/arm-storage/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the encryption scopes available under the specified storage account.
 *
 * @summary Lists all the encryption scopes available under the specified storage account.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/StorageAccountEncryptionScopeList.json
 */
async function storageAccountEncryptionScopeList() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "resource-group-name";
  const accountName = "{storage-account-name}";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.encryptionScopes.list(resourceGroupName, accountName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

storageAccountEncryptionScopeList().catch(console.error);
```
