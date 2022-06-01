Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-compute_18.0.0/sdk/compute/arm-compute/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to get the restore point collection.
 *
 * @summary The operation to get the restore point collection.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/restorePointExamples/RestorePointCollection_Get_WithContainedRestorePoints.json
 */
async function getARestorePointCollectionIncludingTheRestorePointsContainedInTheRestorePointCollection() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const restorePointCollectionName = "rpcName";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.restorePointCollections.get(
    resourceGroupName,
    restorePointCollectionName
  );
  console.log(result);
}

getARestorePointCollectionIncludingTheRestorePointsContainedInTheRestorePointCollection().catch(
  console.error
);
```
