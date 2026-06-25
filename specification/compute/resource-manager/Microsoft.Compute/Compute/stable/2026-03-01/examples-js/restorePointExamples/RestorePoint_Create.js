const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to create the restore point. Updating properties of an existing restore point is not allowed
 *
 * @summary the operation to create the restore point. Updating properties of an existing restore point is not allowed
 * x-ms-original-file: 2026-03-01/restorePointExamples/RestorePoint_Create.json
 */
async function createARestorePoint() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.restorePoints.create("myResourceGroup", "rpcName", "rpName", {
    excludeDisks: [
      {
        id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/disk123",
      },
    ],
    instantAccessDurationMinutes: 120,
  });
  console.log(result);
}
