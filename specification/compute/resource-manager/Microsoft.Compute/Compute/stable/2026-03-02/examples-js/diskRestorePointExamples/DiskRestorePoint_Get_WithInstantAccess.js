const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get disk restorePoint resource
 *
 * @summary get disk restorePoint resource
 * x-ms-original-file: 2026-03-02/diskRestorePointExamples/DiskRestorePoint_Get_WithInstantAccess.json
 */
async function getADiskRestorePointResourceWithInstantAccessSnapshotAccessState() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.diskRestorePoint.get(
    "myResourceGroup",
    "rpc",
    "vmrp",
    "TestDisk45ceb03433006d1baee_5c1528-43e2-4c77-9c55-a78bf5a5fc88",
  );
  console.log(result);
}
