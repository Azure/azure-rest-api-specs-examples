Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-compute_17.3.1/sdk/compute/arm-compute/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to perform maintenance on a virtual machine.
 *
 * @summary The operation to perform maintenance on a virtual machine.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachines_PerformMaintenance_MinimumSet_Gen.json
 */
async function virtualMachinesPerformMaintenanceMinimumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const vmName = "aaaaaaaaaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachines.beginPerformMaintenanceAndWait(
    resourceGroupName,
    vmName
  );
  console.log(result);
}

virtualMachinesPerformMaintenanceMinimumSetGen().catch(console.error);
```
