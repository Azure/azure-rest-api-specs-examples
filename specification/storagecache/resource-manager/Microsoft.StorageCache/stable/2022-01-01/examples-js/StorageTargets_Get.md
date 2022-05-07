Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storagecache_5.1.0/sdk/storagecache/arm-storagecache/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { StorageCacheManagementClient } = require("@azure/arm-storagecache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns a Storage Target from a Cache.
 *
 * @summary Returns a Storage Target from a Cache.
 * x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2022-01-01/examples/StorageTargets_Get.json
 */
async function storageTargetsGet() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "scgroup";
  const cacheName = "sc1";
  const storageTargetName = "st1";
  const credential = new DefaultAzureCredential();
  const client = new StorageCacheManagementClient(credential, subscriptionId);
  const result = await client.storageTargets.get(resourceGroupName, cacheName, storageTargetName);
  console.log(result);
}

storageTargetsGet().catch(console.error);
```
