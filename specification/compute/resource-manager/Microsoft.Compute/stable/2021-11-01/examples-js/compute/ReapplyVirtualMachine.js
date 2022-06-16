const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function reapplyTheStateOfAVirtualMachine() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "ResourceGroup";
  const vmName = "VMName";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachines.beginReapplyAndWait(resourceGroupName, vmName);
  console.log(result);
}

reapplyTheStateOfAVirtualMachine().catch(console.error);
