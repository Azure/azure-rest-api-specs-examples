const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to retrieves information about the run-time state of a virtual machine.
 *
 * @summary retrieves information about the run-time state of a virtual machine.
 * x-ms-original-file: 2026-03-01/virtualMachineExamples/VirtualMachine_Get_InstanceView_WithInterconnectBlock.json
 */
async function getInstanceViewOfAVirtualMachineAssociatedWithAnInterconnectBlock() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachines.instanceView("myResourceGroup", "myVM");
  console.log(result);
}
