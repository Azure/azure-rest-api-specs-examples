const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to delete the run command.
 *
 * @summary The operation to delete the run command.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/runCommandExamples/VirtualMachineRunCommand_Delete.json
 */
async function deleteARunCommand() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const vmName = "myVM";
  const runCommandName = "myRunCommand";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineRunCommands.beginDeleteAndWait(
    resourceGroupName,
    vmName,
    runCommandName
  );
  console.log(result);
}

deleteARunCommand().catch(console.error);
