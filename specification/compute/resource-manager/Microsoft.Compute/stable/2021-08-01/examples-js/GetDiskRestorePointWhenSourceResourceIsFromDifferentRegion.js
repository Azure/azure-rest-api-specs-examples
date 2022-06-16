const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function getAnIncrementalDiskRestorePointWhenSourceResourceIsFromADifferentRegion() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const restorePointCollectionName = "rpc";
  const vmRestorePointName = "vmrp";
  const diskRestorePointName = "TestDisk45ceb03433006d1baee0_b70cd924-3362-4a80-93c2-9415eaa12745";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.diskRestorePointOperations.get(
    resourceGroupName,
    restorePointCollectionName,
    vmRestorePointName,
    diskRestorePointName
  );
  console.log(result);
}

getAnIncrementalDiskRestorePointWhenSourceResourceIsFromADifferentRegion().catch(console.error);
