const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to delete the restore point.
 *
 * @summary The operation to delete the restore point.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/restorePointExamples/RestorePoints_Delete_MaximumSet_Gen.json
 */
async function restorePointsDeleteMaximumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const restorePointCollectionName = "aaaaaaaaaaaaaaaaaaaaaa";
  const restorePointName = "a";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.restorePoints.beginDeleteAndWait(
    resourceGroupName,
    restorePointCollectionName,
    restorePointName
  );
  console.log(result);
}

restorePointsDeleteMaximumSetGen().catch(console.error);
