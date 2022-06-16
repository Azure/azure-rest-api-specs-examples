const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

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
