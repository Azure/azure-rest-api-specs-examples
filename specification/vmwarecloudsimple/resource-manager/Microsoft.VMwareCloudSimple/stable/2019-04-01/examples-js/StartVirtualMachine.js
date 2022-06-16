const { VMwareCloudSimple } = require("@azure/arm-vmwarecloudsimple");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Power on virtual machine
 *
 * @summary Power on virtual machine
 * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/StartVirtualMachine.json
 */
async function startVirtualMachine() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const referer = "https://management.azure.com/";
  const virtualMachineName = "myVirtualMachine";
  const credential = new DefaultAzureCredential();
  const client = new VMwareCloudSimple(credential, subscriptionId);
  const result = await client.virtualMachines.beginStartAndWait(
    resourceGroupName,
    referer,
    virtualMachineName
  );
  console.log(result);
}

startVirtualMachine().catch(console.error);
