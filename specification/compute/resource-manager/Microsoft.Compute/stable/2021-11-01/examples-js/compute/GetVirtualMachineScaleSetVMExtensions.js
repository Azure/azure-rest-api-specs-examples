const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function getVirtualMachineScaleSetVMExtension() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const vmScaleSetName = "myvmScaleSet";
  const instanceId = "0";
  const vmExtensionName = "myVMExtension";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSetVMExtensions.get(
    resourceGroupName,
    vmScaleSetName,
    instanceId,
    vmExtensionName
  );
  console.log(result);
}

getVirtualMachineScaleSetVMExtension().catch(console.error);
