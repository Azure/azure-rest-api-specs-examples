const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to retrieves information about the model view or the instance view of a virtual machine.
 *
 * @summary retrieves information about the model view or the instance view of a virtual machine.
 * x-ms-original-file: 2026-03-01/virtualMachineExamples/VirtualMachine_Get_WithDiskControllerType.json
 */
async function getAVirtualMachineWithDiskControllerTypeProperties() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachines.get("myResourceGroup", "myVM", {
    expand: "userData",
  });
  console.log(result);
}
