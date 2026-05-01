const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets information about a disk access resource.
 *
 * @summary gets information about a disk access resource.
 * x-ms-original-file: 2025-01-02/diskAccessExamples/DiskAccess_Get.json
 */
async function getInformationAboutADiskAccessResource() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.diskAccesses.get("myResourceGroup", "myDiskAccess");
  console.log(result);
}
