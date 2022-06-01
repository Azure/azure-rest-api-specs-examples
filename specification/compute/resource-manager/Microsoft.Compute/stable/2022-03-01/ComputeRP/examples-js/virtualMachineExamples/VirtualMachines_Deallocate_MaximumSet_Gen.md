Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-compute_18.0.0/sdk/compute/arm-compute/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Shuts down the virtual machine and releases the compute resources. You are not billed for the compute resources that this virtual machine uses.
 *
 * @summary Shuts down the virtual machine and releases the compute resources. You are not billed for the compute resources that this virtual machine uses.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineExamples/VirtualMachines_Deallocate_MaximumSet_Gen.json
 */
async function virtualMachinesDeallocateMaximumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const vmName = "aaaaaaaaaa";
  const hibernate = true;
  const options = { hibernate };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachines.beginDeallocateAndWait(
    resourceGroupName,
    vmName,
    options
  );
  console.log(result);
}

virtualMachinesDeallocateMaximumSetGen().catch(console.error);
```
