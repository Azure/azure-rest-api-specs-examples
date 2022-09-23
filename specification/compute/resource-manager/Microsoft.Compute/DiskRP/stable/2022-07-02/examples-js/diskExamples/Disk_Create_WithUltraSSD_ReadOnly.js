const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a disk.
 *
 * @summary Creates or updates a disk.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2022-07-02/examples/diskExamples/Disk_Create_WithUltraSSD_ReadOnly.json
 */
async function createAManagedDiskWithUltraAccountTypeWithReadOnlyPropertySet() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const diskName = "myUltraReadOnlyDisk";
  const disk = {
    creationData: { createOption: "Empty", logicalSectorSize: 4096 },
    diskIopsReadWrite: 125,
    diskMBpsReadWrite: 3000,
    diskSizeGB: 200,
    encryption: { type: "EncryptionAtRestWithPlatformKey" },
    location: "West US",
    sku: { name: "UltraSSD_LRS" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.disks.beginCreateOrUpdateAndWait(resourceGroupName, diskName, disk);
  console.log(result);
}

createAManagedDiskWithUltraAccountTypeWithReadOnlyPropertySet().catch(console.error);
