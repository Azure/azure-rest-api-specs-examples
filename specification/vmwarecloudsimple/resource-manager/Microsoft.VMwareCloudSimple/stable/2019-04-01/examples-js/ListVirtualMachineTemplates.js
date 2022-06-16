const { VMwareCloudSimple } = require("@azure/arm-vmwarecloudsimple");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns list of virtual machine templates in region for private cloud
 *
 * @summary Returns list of virtual machine templates in region for private cloud
 * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/ListVirtualMachineTemplates.json
 */
async function listVirtualMachineTemplates() {
  const subscriptionId = "{subscription-id}";
  const pcName = "myPrivateCloud";
  const regionId = "westus2";
  const resourcePoolName =
    "/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud/resourcePools/resgroup-26";
  const credential = new DefaultAzureCredential();
  const client = new VMwareCloudSimple(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.virtualMachineTemplates.list(pcName, regionId, resourcePoolName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listVirtualMachineTemplates().catch(console.error);
