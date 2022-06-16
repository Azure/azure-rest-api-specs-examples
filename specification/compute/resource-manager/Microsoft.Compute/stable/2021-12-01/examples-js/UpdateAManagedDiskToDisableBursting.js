const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates (patches) a disk.
 *
 * @summary Updates (patches) a disk.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-12-01/examples/UpdateAManagedDiskToDisableBursting.json
 */
async function updateAManagedDiskToDisableBursting() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const diskName = "myDisk";
  const disk = { burstingEnabled: false };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.disks.beginUpdateAndWait(resourceGroupName, diskName, disk);
  console.log(result);
}

updateAManagedDiskToDisableBursting().catch(console.error);
