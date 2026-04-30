const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a disk.
 *
 * @summary creates or updates a disk.
 * x-ms-original-file: 2025-01-02/diskExamples/Disk_Create_WithLogicalSectorSize.json
 */
async function createAnUltraManagedDiskWithLogicalSectorSize512E() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.disks.createOrUpdate("myResourceGroup", "myDisk", {
    location: "West US",
    sku: { name: "UltraSSD_LRS" },
    creationData: { createOption: "Empty", logicalSectorSize: 512 },
    diskSizeGB: 200,
  });
  console.log(result);
}
