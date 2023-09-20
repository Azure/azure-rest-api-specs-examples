const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a disk.
 *
 * @summary Creates or updates a disk.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2023-04-02/examples/diskExamples/Disk_Create_WithPremiumV2_LRSAccountType.json
 */
async function createAManagedDiskWithPremiumV2AccountType() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "myResourceGroup";
  const diskName = "myPremiumV2Disk";
  const disk = {
    creationData: { createOption: "Empty" },
    diskIopsReadWrite: 125,
    diskMBpsReadWrite: 3000,
    diskSizeGB: 200,
    location: "West US",
    sku: { name: "PremiumV2_LRS" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.disks.beginCreateOrUpdateAndWait(resourceGroupName, diskName, disk);
  console.log(result);
}
