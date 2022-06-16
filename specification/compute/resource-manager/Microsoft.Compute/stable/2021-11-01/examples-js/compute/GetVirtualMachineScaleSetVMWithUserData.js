const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function getVMScaleSetVMWithUserData() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const vmScaleSetName = "{vmss-name}";
  const instanceId = "0";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSetVMs.get(
    resourceGroupName,
    vmScaleSetName,
    instanceId
  );
  console.log(result);
}

getVMScaleSetVMWithUserData().catch(console.error);
