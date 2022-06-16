const { ContainerInstanceManagementClient } = require("@azure/arm-containerinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates container group tags with specified values.
 *
 * @summary Updates container group tags with specified values.
 * x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2021-10-01/examples/ContainerGroupsUpdate.json
 */
async function containerGroupsUpdate() {
  const subscriptionId = "subid";
  const resourceGroupName = "demoResource";
  const containerGroupName = "demo1";
  const resource = {
    tags: { tag1key: "tag1Value", tag2key: "tag2Value" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerInstanceManagementClient(credential, subscriptionId);
  const result = await client.containerGroups.update(
    resourceGroupName,
    containerGroupName,
    resource
  );
  console.log(result);
}

containerGroupsUpdate().catch(console.error);
