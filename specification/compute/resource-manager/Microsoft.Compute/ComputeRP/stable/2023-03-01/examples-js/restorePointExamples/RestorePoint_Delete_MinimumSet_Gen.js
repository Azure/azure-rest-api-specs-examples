const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to delete the restore point.
 *
 * @summary The operation to delete the restore point.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-03-01/examples/restorePointExamples/RestorePoint_Delete_MinimumSet_Gen.json
 */
async function restorePointDeleteMinimumSetGen() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "rgcompute";
  const restorePointCollectionName = "aaaaaaaaaaaaaaaaa";
  const restorePointName = "aaaaaaaaaaaaaaaaaaaaaaaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.restorePoints.beginDeleteAndWait(
    resourceGroupName,
    restorePointCollectionName,
    restorePointName
  );
  console.log(result);
}
