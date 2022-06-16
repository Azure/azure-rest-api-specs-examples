const { VMwareCloudSimple } = require("@azure/arm-vmwarecloudsimple");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Patches dedicated node properties
 *
 * @summary Patches dedicated node properties
 * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/PatchDedicatedCloudNode.json
 */
async function patchDedicatedCloudNode() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const dedicatedCloudNodeName = "myNode";
  const dedicatedCloudNodeRequest = {
    tags: { myTag: "tagValue" },
  };
  const credential = new DefaultAzureCredential();
  const client = new VMwareCloudSimple(credential, subscriptionId);
  const result = await client.dedicatedCloudNodes.update(
    resourceGroupName,
    dedicatedCloudNodeName,
    dedicatedCloudNodeRequest
  );
  console.log(result);
}

patchDedicatedCloudNode().catch(console.error);
