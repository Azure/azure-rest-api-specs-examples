const { VMwareCloudSimple } = require("@azure/arm-vmwarecloudsimple");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns dedicated cloud node
 *
 * @summary Returns dedicated cloud node
 * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/GetDedicatedCloudNode.json
 */
async function getDedicatedCloudNode() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const dedicatedCloudNodeName = "myNode";
  const credential = new DefaultAzureCredential();
  const client = new VMwareCloudSimple(credential, subscriptionId);
  const result = await client.dedicatedCloudNodes.get(resourceGroupName, dedicatedCloudNodeName);
  console.log(result);
}

getDedicatedCloudNode().catch(console.error);
