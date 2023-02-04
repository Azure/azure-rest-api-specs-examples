const { StorageCacheManagementClient } = require("@azure/arm-storagecache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a Storage Target. This operation is allowed at any time, but if the Cache is down or unhealthy, the actual creation/modification of the Storage Target may be delayed until the Cache is healthy again.
 *
 * @summary Create or update a Storage Target. This operation is allowed at any time, but if the Cache is down or unhealthy, the actual creation/modification of the Storage Target may be delayed until the Cache is healthy again.
 * x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2022-05-01/examples/StorageTargets_CreateOrUpdate_BlobNfs.json
 */
async function storageTargetsCreateOrUpdateBlobNfs() {
  const subscriptionId =
    process.env["STORAGECACHE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["STORAGECACHE_RESOURCE_GROUP"] || "scgroup";
  const cacheName = "sc1";
  const storageTargetName = "st1";
  const storagetarget = {
    blobNfs: {
      target:
        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.Storage/storageAccounts/blofnfs/blobServices/default/containers/blobnfs",
      usageModel: "WRITE_WORKLOAD_15",
    },
    junctions: [{ namespacePath: "/blobnfs" }],
    targetType: "blobNfs",
  };
  const options = { storagetarget };
  const credential = new DefaultAzureCredential();
  const client = new StorageCacheManagementClient(credential, subscriptionId);
  const result = await client.storageTargets.beginCreateOrUpdateAndWait(
    resourceGroupName,
    cacheName,
    storageTargetName,
    options
  );
  console.log(result);
}
