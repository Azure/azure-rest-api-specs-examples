Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storage_17.2.0/sdk/storage/arm-storage/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the storage accounts available under the given resource group. Note that storage keys are not returned; use the ListKeys operation for this.
 *
 * @summary Lists all the storage accounts available under the given resource group. Note that storage keys are not returned; use the ListKeys operation for this.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/StorageAccountListByResourceGroup.json
 */
async function storageAccountListByResourceGroup() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res6117";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.storageAccounts.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

storageAccountListByResourceGroup().catch(console.error);
```
