const { StorageCacheManagementClient } = require("@azure/arm-storagecache");
const { DefaultAzureCredential } = require("@azure/identity");

async function storageTargetsCreateOrUpdateNoJunctions() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "scgroup";
  const cacheName = "sc1";
  const storageTargetName = "st1";
  const storagetarget = {
    nfs3: { target: "10.0.44.44", usageModel: "READ_HEAVY_INFREQ" },
    targetType: "nfs3",
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

storageTargetsCreateOrUpdateNoJunctions().catch(console.error);
