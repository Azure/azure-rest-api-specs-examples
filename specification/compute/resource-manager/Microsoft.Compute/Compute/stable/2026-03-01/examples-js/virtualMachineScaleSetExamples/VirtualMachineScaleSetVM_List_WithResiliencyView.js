const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets a list of all virtual machines in a VM scale sets.
 *
 * @summary gets a list of all virtual machines in a VM scale sets.
 * x-ms-original-file: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_List_WithResiliencyView.json
 */
async function listVmssVMsWithResilientVMDeletionStatus() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.virtualMachineScaleSetVMs.list("resourceGroupname", "vmssName")) {
    resArray.push(item);
  }

  console.log(resArray);
}
