Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storagecache_5.1.0/sdk/storagecache/arm-storagecache/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { StorageCacheManagementClient } = require("@azure/arm-storagecache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the list of StorageCache.Cache SKUs available to this subscription.
 *
 * @summary Get the list of StorageCache.Cache SKUs available to this subscription.
 * x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2022-01-01/examples/Skus_List.json
 */
async function skusList() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new StorageCacheManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.skus.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

skusList().catch(console.error);
```
