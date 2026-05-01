const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to delete the restore point.
 *
 * @summary the operation to delete the restore point.
 * x-ms-original-file: 2025-11-01/restorePointExamples/RestorePoint_Delete_MinimumSet_Gen.json
 */
async function restorePointDeleteMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  await client.restorePoints.delete("rgcompute", "aaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaa");
}
