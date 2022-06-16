const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function getARestorePointCollectionButNotTheRestorePointsContainedInTheRestorePointCollection() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const restorePointCollectionName = "myRpc";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.restorePointCollections.get(
    resourceGroupName,
    restorePointCollectionName
  );
  console.log(result);
}

getARestorePointCollectionButNotTheRestorePointsContainedInTheRestorePointCollection().catch(
  console.error
);
