const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

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
