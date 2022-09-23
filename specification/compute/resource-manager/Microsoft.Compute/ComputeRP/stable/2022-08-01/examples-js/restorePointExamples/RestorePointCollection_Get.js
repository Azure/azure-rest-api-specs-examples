const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to get the restore point collection.
 *
 * @summary The operation to get the restore point collection.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/restorePointExamples/RestorePointCollection_Get.json
 */
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
