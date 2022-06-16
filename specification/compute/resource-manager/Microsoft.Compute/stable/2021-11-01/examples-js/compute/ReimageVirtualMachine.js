const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function reimageAVirtualMachine() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const vmName = "myVMName";
  const parameters = { tempDisk: true };
  const options = { parameters: parameters };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachines.beginReimageAndWait(
    resourceGroupName,
    vmName,
    options
  );
  console.log(result);
}

reimageAVirtualMachine().catch(console.error);
