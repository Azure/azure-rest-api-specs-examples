const { StorageCacheManagementClient } = require("@azure/arm-storagecache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to suspends client access to a storage target.
 *
 * @summary suspends client access to a storage target.
 * x-ms-original-file: 2026-01-01/StorageTargets_Suspend.json
 */
async function storageTargetsSuspend() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new StorageCacheManagementClient(credential, subscriptionId);
  await client.storageTargetOperations.suspend("scgroup", "sc", "st1");
}
