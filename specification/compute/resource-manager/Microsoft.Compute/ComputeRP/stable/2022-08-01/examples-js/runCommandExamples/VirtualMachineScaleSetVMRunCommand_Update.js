const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to update the VMSS VM run command.
 *
 * @summary The operation to update the VMSS VM run command.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/runCommandExamples/VirtualMachineScaleSetVMRunCommand_Update.json
 */
async function updateVirtualMachineScaleSetVMRunCommand() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const vmScaleSetName = "myvmScaleSet";
  const instanceId = "0";
  const runCommandName = "myRunCommand";
  const runCommand = {
    source: { script: "Write-Host Script Source Updated!" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSetVMRunCommands.beginUpdateAndWait(
    resourceGroupName,
    vmScaleSetName,
    instanceId,
    runCommandName,
    runCommand
  );
  console.log(result);
}

updateVirtualMachineScaleSetVMRunCommand().catch(console.error);
