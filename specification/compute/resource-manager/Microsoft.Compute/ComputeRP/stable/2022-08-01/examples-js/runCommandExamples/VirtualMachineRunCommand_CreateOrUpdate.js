const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to create or update the run command.
 *
 * @summary The operation to create or update the run command.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/runCommandExamples/VirtualMachineRunCommand_CreateOrUpdate.json
 */
async function createOrUpdateARunCommand() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const vmName = "myVM";
  const runCommandName = "myRunCommand";
  const runCommand = {
    asyncExecution: false,
    location: "West US",
    parameters: [
      { name: "param1", value: "value1" },
      { name: "param2", value: "value2" },
    ],
    runAsPassword: "<runAsPassword>",
    runAsUser: "user1",
    source: { script: "Write-Host Hello World!" },
    timeoutInSeconds: 3600,
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineRunCommands.beginCreateOrUpdateAndWait(
    resourceGroupName,
    vmName,
    runCommandName,
    runCommand
  );
  console.log(result);
}

createOrUpdateARunCommand().catch(console.error);
