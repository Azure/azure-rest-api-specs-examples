const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function grantsAccessToADiskRestorePoint() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const restorePointCollectionName = "rpc";
  const vmRestorePointName = "vmrp";
  const diskRestorePointName = "TestDisk45ceb03433006d1baee0_b70cd924-3362-4a80-93c2-9415eaa12745";
  const grantAccessData = {
    access: "Read",
    durationInSeconds: 300,
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.diskRestorePointOperations.beginGrantAccessAndWait(
    resourceGroupName,
    restorePointCollectionName,
    vmRestorePointName,
    diskRestorePointName,
    grantAccessData
  );
  console.log(result);
}

grantsAccessToADiskRestorePoint().catch(console.error);
