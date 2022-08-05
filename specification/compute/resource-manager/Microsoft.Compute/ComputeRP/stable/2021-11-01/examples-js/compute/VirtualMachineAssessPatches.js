const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function assessPatchStateOfAVirtualMachine() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroupName";
  const vmName = "myVMName";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachines.beginAssessPatchesAndWait(resourceGroupName, vmName);
  console.log(result);
}

assessPatchStateOfAVirtualMachine().catch(console.error);
