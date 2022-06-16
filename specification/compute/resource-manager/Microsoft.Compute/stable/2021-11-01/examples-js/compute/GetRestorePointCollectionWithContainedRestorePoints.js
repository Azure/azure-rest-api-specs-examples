const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function getARestorePointCollectionIncludingTheRestorePointsContainedInTheRestorePointCollection() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const restorePointCollectionName = "rpcName";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.restorePointCollections.get(
    resourceGroupName,
    restorePointCollectionName
  );
  console.log(result);
}

getARestorePointCollectionIncludingTheRestorePointsContainedInTheRestorePointCollection().catch(
  console.error
);
