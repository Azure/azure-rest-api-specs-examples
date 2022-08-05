const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function getAnIncrementalDiskRestorePointResource() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const restorePointCollectionName = "rpc";
  const vmRestorePointName = "vmrp";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.diskRestorePointOperations.listByRestorePoint(
    resourceGroupName,
    restorePointCollectionName,
    vmRestorePointName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

getAnIncrementalDiskRestorePointResource().catch(console.error);
