const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to delete the restore point collection. This operation will also delete all the contained restore points.
 *
 * @summary the operation to delete the restore point collection. This operation will also delete all the contained restore points.
 * x-ms-original-file: 2026-03-01/restorePointExamples/RestorePointCollection_Delete_MaximumSet_Gen.json
 */
async function restorePointCollectionDeleteMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  await client.restorePointCollections.delete("rgcompute", "aaaaaaaaaaaaaaaaa");
}
