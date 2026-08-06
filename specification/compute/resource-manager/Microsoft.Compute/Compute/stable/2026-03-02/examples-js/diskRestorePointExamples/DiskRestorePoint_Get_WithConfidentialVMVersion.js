const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get disk restorePoint resource
 *
 * @summary get disk restorePoint resource
 * x-ms-original-file: 2026-03-02/diskRestorePointExamples/DiskRestorePoint_Get_WithConfidentialVMVersion.json
 */
async function getAConfidentialVMIncrementalDiskRestorePointResourceWithConfidentialVMVersion() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.diskRestorePoint.get(
    "myResourceGroup",
    "rpc",
    "vmrp",
    "myConfidentialDisk_c4bc27e0-ccf6-494e-a740-af34de775527",
  );
  console.log(result);
}
