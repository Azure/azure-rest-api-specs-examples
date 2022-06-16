const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function getAVirtualMachineScaleSetWithUserData() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const vmScaleSetName = "myVirtualMachineScaleSet";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSets.get(resourceGroupName, vmScaleSetName);
  console.log(result);
}

getAVirtualMachineScaleSetWithUserData().catch(console.error);
