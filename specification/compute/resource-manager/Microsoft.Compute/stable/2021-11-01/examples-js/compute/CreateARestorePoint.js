const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to create the restore point. Updating properties of an existing restore point is not allowed
 *
 * @summary The operation to create the restore point. Updating properties of an existing restore point is not allowed
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/CreateARestorePoint.json
 */
async function createARestorePoint() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const restorePointCollectionName = "rpcName";
  const restorePointName = "rpName";
  const parameters = {
    excludeDisks: [
      {
        id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/disk123",
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.restorePoints.beginCreateAndWait(
    resourceGroupName,
    restorePointCollectionName,
    restorePointName,
    parameters
  );
  console.log(result);
}

createARestorePoint().catch(console.error);
