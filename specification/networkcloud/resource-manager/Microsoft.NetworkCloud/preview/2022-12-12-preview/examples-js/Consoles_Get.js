const { NetworkCloud } = require("@azure/arm-networkcloud");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get properties of the provided virtual machine console.
 *
 * @summary Get properties of the provided virtual machine console.
 * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/Consoles_Get.json
 */
async function getVirtualMachineConsole() {
  const subscriptionId = process.env["NETWORKCLOUD_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName = process.env["NETWORKCLOUD_RESOURCE_GROUP"] || "resourceGroupName";
  const virtualMachineName = "virtualMachineName";
  const consoleName = "default";
  const credential = new DefaultAzureCredential();
  const client = new NetworkCloud(credential, subscriptionId);
  const result = await client.consoles.get(resourceGroupName, virtualMachineName, consoleName);
  console.log(result);
}
