const { VMwareCloudSimple } = require("@azure/arm-vmwarecloudsimple");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete dedicated cloud node
 *
 * @summary Delete dedicated cloud node
 * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/DeleteDedicatedCloudNode.json
 */
async function deleteDedicatedCloudNode() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const dedicatedCloudNodeName = "myNode";
  const credential = new DefaultAzureCredential();
  const client = new VMwareCloudSimple(credential, subscriptionId);
  const result = await client.dedicatedCloudNodes.delete(resourceGroupName, dedicatedCloudNodeName);
  console.log(result);
}

deleteDedicatedCloudNode().catch(console.error);
