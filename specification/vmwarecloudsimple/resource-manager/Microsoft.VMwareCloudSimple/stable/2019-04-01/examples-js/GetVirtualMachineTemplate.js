const { VMwareCloudSimple } = require("@azure/arm-vmwarecloudsimple");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns virtual machine templates by its name
 *
 * @summary Returns virtual machine templates by its name
 * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/GetVirtualMachineTemplate.json
 */
async function getVirtualMachineTemplate() {
  const subscriptionId = process.env["VMWARECLOUDSIMPLE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const regionId = "westus2";
  const pcName = "myPrivateCloud";
  const virtualMachineTemplateName = "vm-34";
  const credential = new DefaultAzureCredential();
  const client = new VMwareCloudSimple(credential, subscriptionId);
  const result = await client.virtualMachineTemplates.get(
    regionId,
    pcName,
    virtualMachineTemplateName
  );
  console.log(result);
}
