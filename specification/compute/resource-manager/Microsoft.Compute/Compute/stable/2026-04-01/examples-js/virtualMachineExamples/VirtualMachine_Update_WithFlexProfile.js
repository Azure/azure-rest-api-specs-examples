const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to update a virtual machine.
 *
 * @summary the operation to update a virtual machine.
 * x-ms-original-file: 2026-04-01/virtualMachineExamples/VirtualMachine_Update_WithFlexProfile.json
 */
async function updateAVMFlexProfile() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachines.update("myResourceGroup", "myVM", {
    flexProfile: { preferredVMSize: "Standard_D8s_v5" },
  });
  console.log(result);
}
