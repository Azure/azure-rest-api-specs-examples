Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storage_17.2.0/sdk/storage/arm-storage/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the blob inventory policy associated with the specified storage account.
 *
 * @summary Gets the blob inventory policy associated with the specified storage account.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/StorageAccountListBlobInventoryPolicy.json
 */
async function storageAccountGetBlobInventoryPolicy() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res7687";
  const accountName = "sto9699";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.blobInventoryPolicies.list(resourceGroupName, accountName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

storageAccountGetBlobInventoryPolicy().catch(console.error);
```
