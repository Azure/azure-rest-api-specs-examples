Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storage_17.2.0/sdk/storage/arm-storage/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the access keys or Kerberos keys (if active directory enabled) for the specified storage account.
 *
 * @summary Lists the access keys or Kerberos keys (if active directory enabled) for the specified storage account.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/StorageAccountListKeys.json
 */
async function storageAccountListKeys() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res418";
  const accountName = "sto2220";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.storageAccounts.listKeys(resourceGroupName, accountName);
  console.log(result);
}

storageAccountListKeys().catch(console.error);
```
