const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a disk.
 *
 * @summary creates or updates a disk.
 * x-ms-original-file: 2025-01-02/diskExamples/Disk_Create_WithUltraSSD_ReadOnly.json
 */
async function createAManagedDiskWithUltraAccountTypeWithReadOnlyPropertySet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.disks.createOrUpdate("myResourceGroup", "myUltraReadOnlyDisk", {
    location: "West US",
    sku: { name: "UltraSSD_LRS" },
    creationData: { createOption: "Empty", logicalSectorSize: 4096 },
    diskSizeGB: 200,
    diskIopsReadWrite: 125,
    diskMBpsReadWrite: 3000,
    encryption: { type: "EncryptionAtRestWithPlatformKey" },
  });
  console.log(result);
}
