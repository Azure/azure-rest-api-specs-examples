const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves information about the run-time state of a virtual machine.
 *
 * @summary Retrieves information about the run-time state of a virtual machine.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineExamples/VirtualMachine_Get_InstanceView.json
 */
async function getVirtualMachineInstanceView() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const vmName = "myVM";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachines.instanceView(resourceGroupName, vmName);
  console.log(result);
}

getVirtualMachineInstanceView().catch(console.error);
