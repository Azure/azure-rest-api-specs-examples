const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets a virtual machine from a VM scale set.
 *
 * @summary gets a virtual machine from a VM scale set.
 * x-ms-original-file: 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_Get_WithVirtualMachineResourceId.json
 */
async function getVMScaleSetFlexVMWithVirtualMachineResourceId() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSetVMs.get(
    "myResourceGroup",
    "{vmss-flex-name}",
    "{vmss-flex-vm-name}",
  );
  console.log(result);
}
