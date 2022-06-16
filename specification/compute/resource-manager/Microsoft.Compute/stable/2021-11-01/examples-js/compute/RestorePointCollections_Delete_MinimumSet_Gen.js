const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function restorePointCollectionsDeleteMinimumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const restorePointCollectionName = "aaaaaaaaaaaaaaaaaaaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.restorePointCollections.beginDeleteAndWait(
    resourceGroupName,
    restorePointCollectionName
  );
  console.log(result);
}

restorePointCollectionsDeleteMinimumSetGen().catch(console.error);
