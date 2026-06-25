const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to get the restore point.
 *
 * @summary the operation to get the restore point.
 * x-ms-original-file: 2026-03-01/restorePointExamples/RestorePoint_Get_WithInstanceView.json
 */
async function getRestorePointWithInstanceView() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.restorePoints.get("myResourceGroup", "rpcName", "rpName");
  console.log(result);
}
