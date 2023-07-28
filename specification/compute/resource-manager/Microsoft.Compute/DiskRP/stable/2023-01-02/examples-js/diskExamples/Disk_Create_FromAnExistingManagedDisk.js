const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a disk.
 *
 * @summary Creates or updates a disk.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2023-01-02/examples/diskExamples/Disk_Create_FromAnExistingManagedDisk.json
 */
async function createAManagedDiskFromAnExistingManagedDiskInTheSameOrDifferentSubscription() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "myResourceGroup";
  const diskName = "myDisk2";
  const disk = {
    creationData: {
      createOption: "Copy",
      sourceResourceId:
        "subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myDisk1",
    },
    location: "West US",
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.disks.beginCreateOrUpdateAndWait(resourceGroupName, diskName, disk);
  console.log(result);
}
