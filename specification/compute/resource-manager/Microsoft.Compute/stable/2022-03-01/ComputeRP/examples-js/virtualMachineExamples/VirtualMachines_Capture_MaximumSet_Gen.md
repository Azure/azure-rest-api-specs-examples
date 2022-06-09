Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-compute_19.0.0/sdk/compute/arm-compute/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Captures the VM by copying virtual hard disks of the VM and outputs a template that can be used to create similar VMs.
 *
 * @summary Captures the VM by copying virtual hard disks of the VM and outputs a template that can be used to create similar VMs.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineExamples/VirtualMachines_Capture_MaximumSet_Gen.json
 */
async function virtualMachinesCaptureMaximumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const vmName = "aaaaaaaaaaaaaaaaaaaa";
  const parameters = {
    destinationContainerName: "aaaaaaa",
    overwriteVhds: true,
    vhdPrefix: "aaaaaaaaa",
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachines.beginCaptureAndWait(
    resourceGroupName,
    vmName,
    parameters
  );
  console.log(result);
}

virtualMachinesCaptureMaximumSetGen().catch(console.error);
```
