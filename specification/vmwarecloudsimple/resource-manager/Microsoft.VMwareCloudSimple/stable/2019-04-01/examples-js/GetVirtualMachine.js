const { VMwareCloudSimple } = require("@azure/arm-vmwarecloudsimple");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get virtual machine
 *
 * @summary Get virtual machine
 * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/GetVirtualMachine.json
 */
async function getVirtualMachine() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const virtualMachineName = "myVirtualMachine";
  const credential = new DefaultAzureCredential();
  const client = new VMwareCloudSimple(credential, subscriptionId);
  const result = await client.virtualMachines.get(resourceGroupName, virtualMachineName);
  console.log(result);
}

getVirtualMachine().catch(console.error);
