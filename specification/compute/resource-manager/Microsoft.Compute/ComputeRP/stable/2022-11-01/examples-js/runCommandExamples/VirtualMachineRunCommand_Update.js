const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to update the run command.
 *
 * @summary The operation to update the run command.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-11-01/examples/runCommandExamples/VirtualMachineRunCommand_Update.json
 */
async function updateARunCommand() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "myResourceGroup";
  const vmName = "myVM";
  const runCommandName = "myRunCommand";
  const runCommand = {
    source: { script: "Write-Host Script Source Updated!" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineRunCommands.beginUpdateAndWait(
    resourceGroupName,
    vmName,
    runCommandName,
    runCommand
  );
  console.log(result);
}
