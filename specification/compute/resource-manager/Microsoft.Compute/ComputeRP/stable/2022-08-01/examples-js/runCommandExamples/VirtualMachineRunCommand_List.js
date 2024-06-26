const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to get all run commands of a Virtual Machine.
 *
 * @summary The operation to get all run commands of a Virtual Machine.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/runCommandExamples/VirtualMachineRunCommand_List.json
 */
async function listRunCommandsInAVirtualMachine() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const vmName = "myVM";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.virtualMachineRunCommands.listByVirtualMachine(
    resourceGroupName,
    vmName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listRunCommandsInAVirtualMachine().catch(console.error);
