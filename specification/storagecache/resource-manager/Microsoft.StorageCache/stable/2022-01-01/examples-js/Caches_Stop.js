const { StorageCacheManagementClient } = require("@azure/arm-storagecache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Tells an Active Cache to transition to Stopped state.
 *
 * @summary Tells an Active Cache to transition to Stopped state.
 * x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2022-01-01/examples/Caches_Stop.json
 */
async function cachesStop() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "scgroup";
  const cacheName = "sc";
  const credential = new DefaultAzureCredential();
  const client = new StorageCacheManagementClient(credential, subscriptionId);
  const result = await client.caches.beginStopAndWait(resourceGroupName, cacheName);
  console.log(result);
}

cachesStop().catch(console.error);
