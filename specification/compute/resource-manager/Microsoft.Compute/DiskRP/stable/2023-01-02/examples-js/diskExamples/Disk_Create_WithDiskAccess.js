const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a disk.
 *
 * @summary Creates or updates a disk.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2023-01-02/examples/diskExamples/Disk_Create_WithDiskAccess.json
 */
async function createAManagedDiskAndAssociateWithDiskAccessResource() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "myResourceGroup";
  const diskName = "myDisk";
  const disk = {
    creationData: { createOption: "Empty" },
    diskAccessId:
      "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskAccesses/{existing-diskAccess-name}",
    diskSizeGB: 200,
    location: "West US",
    networkAccessPolicy: "AllowPrivate",
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.disks.beginCreateOrUpdateAndWait(resourceGroupName, diskName, disk);
  console.log(result);
}
