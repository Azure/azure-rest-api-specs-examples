const { StorageCacheManagementClient } = require("@azure/arm-storagecache");
const { DefaultAzureCredential } = require("@azure/identity");

async function storageTargetsCreateOrUpdateBlobNfs() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "scgroup";
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
  const options = { storagetarget: storagetarget };
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

storageTargetsCreateOrUpdateBlobNfs().catch(console.error);
