Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-compute_18.0.0/sdk/compute/arm-compute/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Grants access to a diskRestorePoint.
 *
 * @summary Grants access to a diskRestorePoint.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-12-01/examples/BeginGetAccessDiskRestorePoint.json
 */
async function grantsAccessToADiskRestorePoint() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const restorePointCollectionName = "rpc";
  const vmRestorePointName = "vmrp";
  const diskRestorePointName = "TestDisk45ceb03433006d1baee0_b70cd924-3362-4a80-93c2-9415eaa12745";
  const grantAccessData = {
    access: "Read",
    durationInSeconds: 300,
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.diskRestorePointOperations.beginGrantAccessAndWait(
    resourceGroupName,
    restorePointCollectionName,
    vmRestorePointName,
    diskRestorePointName,
    grantAccessData
  );
  console.log(result);
}

grantsAccessToADiskRestorePoint().catch(console.error);
```
