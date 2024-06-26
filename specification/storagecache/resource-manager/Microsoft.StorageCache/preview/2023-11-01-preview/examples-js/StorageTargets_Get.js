const { StorageCacheManagementClient } = require("@azure/arm-storagecache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns a Storage Target from a cache.
 *
 * @summary Returns a Storage Target from a cache.
 * x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/preview/2023-11-01-preview/examples/StorageTargets_Get.json
 */
async function storageTargetsGet() {
  const subscriptionId =
    process.env["STORAGECACHE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["STORAGECACHE_RESOURCE_GROUP"] || "scgroup";
  const cacheName = "sc1";
  const storageTargetName = "st1";
  const credential = new DefaultAzureCredential();
  const client = new StorageCacheManagementClient(credential, subscriptionId);
  const result = await client.storageTargets.get(resourceGroupName, cacheName, storageTargetName);
  console.log(result);
}
