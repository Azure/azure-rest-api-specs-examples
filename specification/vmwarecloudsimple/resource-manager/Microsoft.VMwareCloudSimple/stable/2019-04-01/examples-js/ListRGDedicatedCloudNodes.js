const { VMwareCloudSimple } = require("@azure/arm-vmwarecloudsimple");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns list of dedicate cloud nodes within resource group
 *
 * @summary Returns list of dedicate cloud nodes within resource group
 * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/ListRGDedicatedCloudNodes.json
 */
async function listRgDedicatedCloudNodes() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const credential = new DefaultAzureCredential();
  const client = new VMwareCloudSimple(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.dedicatedCloudNodes.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listRgDedicatedCloudNodes().catch(console.error);
