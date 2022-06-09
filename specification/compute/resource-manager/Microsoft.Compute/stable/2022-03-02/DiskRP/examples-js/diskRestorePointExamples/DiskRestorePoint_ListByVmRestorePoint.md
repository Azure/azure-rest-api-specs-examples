Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-compute_19.0.0/sdk/compute/arm-compute/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists diskRestorePoints under a vmRestorePoint.
 *
 * @summary Lists diskRestorePoints under a vmRestorePoint.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-02/DiskRP/examples/diskRestorePointExamples/DiskRestorePoint_ListByVmRestorePoint.json
 */
async function getAnIncrementalDiskRestorePointResource() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const restorePointCollectionName = "rpc";
  const vmRestorePointName = "vmrp";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.diskRestorePointOperations.listByRestorePoint(
    resourceGroupName,
    restorePointCollectionName,
    vmRestorePointName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

getAnIncrementalDiskRestorePointResource().catch(console.error);
```
