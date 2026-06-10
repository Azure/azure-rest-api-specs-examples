const { StorageCacheManagementClient } = require("@azure/arm-storagecache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a Storage Target. This operation is allowed at any time, but if the cache is down or unhealthy, the actual creation/modification of the Storage Target may be delayed until the cache is healthy again.
 *
 * @summary create or update a Storage Target. This operation is allowed at any time, but if the cache is down or unhealthy, the actual creation/modification of the Storage Target may be delayed until the cache is healthy again.
 * x-ms-original-file: 2026-01-01/StorageTargets_CreateOrUpdate_NoJunctions.json
 */
async function storageTargetsCreateOrUpdateNoJunctions() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new StorageCacheManagementClient(credential, subscriptionId);
  const result = await client.storageTargets.createOrUpdate("scgroup", "sc1", "st1", {
    nfs3: { target: "10.0.44.44", usageModel: "READ_ONLY", verificationTimer: 30 },
    targetType: "nfs3",
  });
  console.log(result);
}
