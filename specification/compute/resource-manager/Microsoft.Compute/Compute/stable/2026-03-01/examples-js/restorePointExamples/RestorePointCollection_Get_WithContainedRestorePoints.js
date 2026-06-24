const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to get the restore point collection.
 *
 * @summary the operation to get the restore point collection.
 * x-ms-original-file: 2026-03-01/restorePointExamples/RestorePointCollection_Get_WithContainedRestorePoints.json
 */
async function getARestorePointCollectionIncludingTheRestorePointsContainedInTheRestorePointCollection() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.restorePointCollections.get("myResourceGroup", "rpcName");
  console.log(result);
}
