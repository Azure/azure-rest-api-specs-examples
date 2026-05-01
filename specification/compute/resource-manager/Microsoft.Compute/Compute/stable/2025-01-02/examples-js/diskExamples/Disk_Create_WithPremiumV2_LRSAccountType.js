const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a disk.
 *
 * @summary creates or updates a disk.
 * x-ms-original-file: 2025-01-02/diskExamples/Disk_Create_WithPremiumV2_LRSAccountType.json
 */
async function createAManagedDiskWithPremiumV2AccountType() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.disks.createOrUpdate("myResourceGroup", "myPremiumV2Disk", {
    location: "West US",
    sku: { name: "PremiumV2_LRS" },
    creationData: { createOption: "Empty" },
    diskSizeGB: 200,
    diskIopsReadWrite: 125,
    diskMBpsReadWrite: 3000,
  });
  console.log(result);
}
