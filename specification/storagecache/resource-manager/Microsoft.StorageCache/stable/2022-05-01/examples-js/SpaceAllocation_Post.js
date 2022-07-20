const { StorageCacheManagementClient } = require("@azure/arm-storagecache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update cache space allocation.
 *
 * @summary Update cache space allocation.
 * x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2022-05-01/examples/SpaceAllocation_Post.json
 */
async function spaceAllocationPost() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "scgroup";
  const cacheName = "sc1";
  const spaceAllocation = [
    { name: "st1", allocationPercentage: 25 },
    { name: "st2", allocationPercentage: 50 },
    { name: "st3", allocationPercentage: 25 },
  ];
  const options = { spaceAllocation };
  const credential = new DefaultAzureCredential();
  const client = new StorageCacheManagementClient(credential, subscriptionId);
  const result = await client.caches.beginSpaceAllocationAndWait(
    resourceGroupName,
    cacheName,
    options
  );
  console.log(result);
}

spaceAllocationPost().catch(console.error);
