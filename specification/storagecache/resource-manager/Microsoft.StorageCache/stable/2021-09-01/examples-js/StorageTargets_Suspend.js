const { StorageCacheManagementClient } = require("@azure/arm-storagecache");
const { DefaultAzureCredential } = require("@azure/identity");

async function storageTargetsSuspend() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "scgroup";
  const cacheName = "sc";
  const storageTargetName = "st1";
  const credential = new DefaultAzureCredential();
  const client = new StorageCacheManagementClient(credential, subscriptionId);
  const result = await client.storageTargetOperations.beginSuspendAndWait(
    resourceGroupName,
    cacheName,
    storageTargetName
  );
  console.log(result);
}

storageTargetsSuspend().catch(console.error);
