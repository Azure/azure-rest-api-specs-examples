const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to grants access to a disk.
 *
 * @summary grants access to a disk.
 * x-ms-original-file: 2025-01-02/diskExamples/Disk_BeginGetAccess.json
 */
async function getASasOnAManagedDisk() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.disks.grantAccess("myResourceGroup", "myDisk", {
    access: "Read",
    durationInSeconds: 300,
    fileFormat: "VHD",
  });
  console.log(result);
}
