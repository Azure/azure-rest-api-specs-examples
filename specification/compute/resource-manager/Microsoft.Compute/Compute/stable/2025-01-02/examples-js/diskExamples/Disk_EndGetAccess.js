const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to revokes access to a disk.
 *
 * @summary revokes access to a disk.
 * x-ms-original-file: 2025-01-02/diskExamples/Disk_EndGetAccess.json
 */
async function revokeAccessToAManagedDisk() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  await client.disks.revokeAccess("myResourceGroup", "myDisk");
}
