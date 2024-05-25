const { StorageCacheManagementClient } = require("@azure/arm-storagecache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a Storage Target. This operation is allowed at any time, but if the cache is down or unhealthy, the actual creation/modification of the Storage Target may be delayed until the cache is healthy again.
 *
 * @summary Create or update a Storage Target. This operation is allowed at any time, but if the cache is down or unhealthy, the actual creation/modification of the Storage Target may be delayed until the cache is healthy again.
 * x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2024-03-01/examples/StorageTargets_CreateOrUpdate.json
 */
async function storageTargetsCreateOrUpdate() {
  const subscriptionId =
    process.env["STORAGECACHE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["STORAGECACHE_RESOURCE_GROUP"] || "scgroup";
  const cacheName = "sc1";
  const storageTargetName = "st1";
  const storagetarget = {
    junctions: [
      {
        namespacePath: "/path/on/cache",
        nfsAccessPolicy: "default",
        nfsExport: "exp1",
        targetPath: "/path/on/exp1",
      },
      {
        namespacePath: "/path2/on/cache",
        nfsAccessPolicy: "rootSquash",
        nfsExport: "exp2",
        targetPath: "/path2/on/exp2",
      },
    ],
    nfs3: {
      target: "10.0.44.44",
      usageModel: "READ_ONLY",
      verificationTimer: 30,
    },
    targetType: "nfs3",
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageCacheManagementClient(credential, subscriptionId);
  const result = await client.storageTargets.beginCreateOrUpdateAndWait(
    resourceGroupName,
    cacheName,
    storageTargetName,
    storagetarget,
  );
  console.log(result);
}
