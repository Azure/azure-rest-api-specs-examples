const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function getARestorePoint() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const restorePointCollectionName = "rpcName";
  const restorePointName = "rpName";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.restorePoints.get(
    resourceGroupName,
    restorePointCollectionName,
    restorePointName
  );
  console.log(result);
}

getARestorePoint().catch(console.error);
