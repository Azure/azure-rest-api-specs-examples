const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to delete the restore point collection. This operation will also delete all the contained restore points.
 *
 * @summary The operation to delete the restore point collection. This operation will also delete all the contained restore points.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/RestorePointCollections_Delete_MaximumSet_Gen.json
 */
async function restorePointCollectionsDeleteMaximumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const restorePointCollectionName = "aaaaaaaaaaaaaaaaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.restorePointCollections.beginDeleteAndWait(
    resourceGroupName,
    restorePointCollectionName
  );
  console.log(result);
}

restorePointCollectionsDeleteMaximumSetGen().catch(console.error);
