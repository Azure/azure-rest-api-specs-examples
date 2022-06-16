const { VMwareCloudSimple } = require("@azure/arm-vmwarecloudsimple");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns list of dedicated cloud services within a resource group
 *
 * @summary Returns list of dedicated cloud services within a resource group
 * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/ListRGDedicatedCloudServices.json
 */
async function listRgDedicatedCloudServices() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const credential = new DefaultAzureCredential();
  const client = new VMwareCloudSimple(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.dedicatedCloudServices.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listRgDedicatedCloudServices().catch(console.error);
