const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all available virtual machine sizes to which the specified virtual machine can be resized.
 *
 * @summary Lists all available virtual machine sizes to which the specified virtual machine can be resized.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineExamples/VirtualMachines_ListAvailableVmSizes.json
 */
async function listsAllAvailableVirtualMachineSizesToWhichTheSpecifiedVirtualMachineCanBeResized() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const vmName = "myVmName";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.virtualMachines.listAvailableSizes(resourceGroupName, vmName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listsAllAvailableVirtualMachineSizesToWhichTheSpecifiedVirtualMachineCanBeResized().catch(
  console.error
);
