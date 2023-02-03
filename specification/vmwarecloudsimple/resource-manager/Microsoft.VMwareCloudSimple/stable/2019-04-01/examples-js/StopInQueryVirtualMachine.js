const { VMwareCloudSimple } = require("@azure/arm-vmwarecloudsimple");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Power off virtual machine, options: shutdown, poweroff, and suspend
 *
 * @summary Power off virtual machine, options: shutdown, poweroff, and suspend
 * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/StopInQueryVirtualMachine.json
 */
async function stopInQueryVirtualMachine() {
  const subscriptionId = process.env["VMWARECLOUDSIMPLE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["VMWARECLOUDSIMPLE_RESOURCE_GROUP"] || "myResourceGroup";
  const referer = "https://management.azure.com/";
  const virtualMachineName = "myVirtualMachine";
  const mode = "suspend";
  const options = { mode };
  const credential = new DefaultAzureCredential();
  const client = new VMwareCloudSimple(credential, subscriptionId);
  const result = await client.virtualMachines.beginStopAndWait(
    resourceGroupName,
    referer,
    virtualMachineName,
    options
  );
  console.log(result);
}
