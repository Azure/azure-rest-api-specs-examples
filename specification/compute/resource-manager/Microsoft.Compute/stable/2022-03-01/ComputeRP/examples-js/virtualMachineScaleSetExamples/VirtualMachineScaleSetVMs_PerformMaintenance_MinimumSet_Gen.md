Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-compute_18.0.0/sdk/compute/arm-compute/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Performs maintenance on a virtual machine in a VM scale set.
 *
 * @summary Performs maintenance on a virtual machine in a VM scale set.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetVMs_PerformMaintenance_MinimumSet_Gen.json
 */
async function virtualMachineScaleSetVMSPerformMaintenanceMinimumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const vmScaleSetName = "aaaaaaaaaa";
  const instanceId = "aaaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSetVMs.beginPerformMaintenanceAndWait(
    resourceGroupName,
    vmScaleSetName,
    instanceId
  );
  console.log(result);
}

virtualMachineScaleSetVMSPerformMaintenanceMinimumSetGen().catch(console.error);
```
