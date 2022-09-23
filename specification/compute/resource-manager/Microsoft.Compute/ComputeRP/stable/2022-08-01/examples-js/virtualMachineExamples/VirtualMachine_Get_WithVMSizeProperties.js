const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves information about the model view or the instance view of a virtual machine.
 *
 * @summary Retrieves information about the model view or the instance view of a virtual machine.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/virtualMachineExamples/VirtualMachine_Get_WithVMSizeProperties.json
 */
async function getAVirtualMachineWithVMSizeProperties() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const vmName = "myVM";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachines.get(resourceGroupName, vmName);
  console.log(result);
}

getAVirtualMachineWithVMSizeProperties().catch(console.error);
