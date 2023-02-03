const { VMwareCloudSimple } = require("@azure/arm-vmwarecloudsimple");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete virtual machine
 *
 * @summary Delete virtual machine
 * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/DeleteVirtualMachine.json
 */
async function deleteVirtualMachine() {
  const subscriptionId = process.env["VMWARECLOUDSIMPLE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["VMWARECLOUDSIMPLE_RESOURCE_GROUP"] || "myResourceGroup";
  const referer = "https://management.azure.com/";
  const virtualMachineName = "myVirtualMachine";
  const credential = new DefaultAzureCredential();
  const client = new VMwareCloudSimple(credential, subscriptionId);
  const result = await client.virtualMachines.beginDeleteAndWait(
    resourceGroupName,
    referer,
    virtualMachineName
  );
  console.log(result);
}
