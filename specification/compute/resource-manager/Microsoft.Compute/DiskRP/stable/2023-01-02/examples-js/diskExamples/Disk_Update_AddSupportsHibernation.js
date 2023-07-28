const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates (patches) a disk.
 *
 * @summary Updates (patches) a disk.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2023-01-02/examples/diskExamples/Disk_Update_AddSupportsHibernation.json
 */
async function updateAManagedDiskToAddSupportsHibernation() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "myResourceGroup";
  const diskName = "myDisk";
  const disk = { supportsHibernation: true };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.disks.beginUpdateAndWait(resourceGroupName, diskName, disk);
  console.log(result);
}
