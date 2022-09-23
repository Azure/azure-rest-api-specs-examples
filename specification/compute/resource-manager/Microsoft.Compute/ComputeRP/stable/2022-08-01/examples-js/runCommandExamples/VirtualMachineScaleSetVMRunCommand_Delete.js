const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to delete the VMSS VM run command.
 *
 * @summary The operation to delete the VMSS VM run command.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/runCommandExamples/VirtualMachineScaleSetVMRunCommand_Delete.json
 */
async function deleteVirtualMachineScaleSetVMRunCommand() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const vmScaleSetName = "myvmScaleSet";
  const instanceId = "0";
  const runCommandName = "myRunCommand";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSetVMRunCommands.beginDeleteAndWait(
    resourceGroupName,
    vmScaleSetName,
    instanceId,
    runCommandName
  );
  console.log(result);
}

deleteVirtualMachineScaleSetVMRunCommand().catch(console.error);
