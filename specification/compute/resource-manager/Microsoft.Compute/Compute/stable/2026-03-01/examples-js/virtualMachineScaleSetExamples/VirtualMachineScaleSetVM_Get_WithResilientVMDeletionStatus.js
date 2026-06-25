const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets a virtual machine from a VM scale set.
 *
 * @summary gets a virtual machine from a VM scale set.
 * x-ms-original-file: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_Get_WithResilientVMDeletionStatus.json
 */
async function getVMScaleSetVMWithResiliencyView() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSetVMs.get("myResourceGroup", "{vmss-name}", "1");
  console.log(result);
}
