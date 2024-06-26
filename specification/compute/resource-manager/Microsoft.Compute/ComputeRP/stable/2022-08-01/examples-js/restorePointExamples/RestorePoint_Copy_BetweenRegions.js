const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to create the restore point. Updating properties of an existing restore point is not allowed
 *
 * @summary The operation to create the restore point. Updating properties of an existing restore point is not allowed
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/restorePointExamples/RestorePoint_Copy_BetweenRegions.json
 */
async function copyARestorePointToADifferentRegion() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const restorePointCollectionName = "rpcName";
  const restorePointName = "rpName";
  const parameters = {
    sourceRestorePoint: {
      id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/sourceRpcName/restorePoints/sourceRpName",
    },
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

copyARestorePointToADifferentRegion().catch(console.error);
