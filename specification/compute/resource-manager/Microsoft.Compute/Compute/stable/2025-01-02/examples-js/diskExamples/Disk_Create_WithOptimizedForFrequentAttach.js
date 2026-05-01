const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a disk.
 *
 * @summary creates or updates a disk.
 * x-ms-original-file: 2025-01-02/diskExamples/Disk_Create_WithOptimizedForFrequentAttach.json
 */
async function createAManagedDiskWithOptimizedForFrequentAttach() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.disks.createOrUpdate("myResourceGroup", "myDisk", {
    location: "West US",
    creationData: { createOption: "Empty" },
    diskSizeGB: 200,
    optimizedForFrequentAttach: true,
  });
  console.log(result);
}
