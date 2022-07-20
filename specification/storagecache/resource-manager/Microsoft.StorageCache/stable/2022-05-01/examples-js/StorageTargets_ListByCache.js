const { StorageCacheManagementClient } = require("@azure/arm-storagecache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns a list of Storage Targets for the specified Cache.
 *
 * @summary Returns a list of Storage Targets for the specified Cache.
 * x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2022-05-01/examples/StorageTargets_ListByCache.json
 */
async function storageTargetsList() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "scgroup";
  const cacheName = "sc1";
  const credential = new DefaultAzureCredential();
  const client = new StorageCacheManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.storageTargets.listByCache(resourceGroupName, cacheName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

storageTargetsList().catch(console.error);
