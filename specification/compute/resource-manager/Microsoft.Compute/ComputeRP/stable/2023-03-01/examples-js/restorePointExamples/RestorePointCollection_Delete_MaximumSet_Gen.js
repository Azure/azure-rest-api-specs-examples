const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to delete the restore point collection. This operation will also delete all the contained restore points.
 *
 * @summary The operation to delete the restore point collection. This operation will also delete all the contained restore points.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-03-01/examples/restorePointExamples/RestorePointCollection_Delete_MaximumSet_Gen.json
 */
async function restorePointCollectionDeleteMaximumSetGen() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "rgcompute";
  const restorePointCollectionName = "aaaaaaaaaaaaaaaaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.restorePointCollections.beginDeleteAndWait(
    resourceGroupName,
    restorePointCollectionName
  );
  console.log(result);
}
