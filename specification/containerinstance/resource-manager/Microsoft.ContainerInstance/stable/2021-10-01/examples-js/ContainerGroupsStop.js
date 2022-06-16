const { ContainerInstanceManagementClient } = require("@azure/arm-containerinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Stops all containers in a container group. Compute resources will be deallocated and billing will stop.
 *
 * @summary Stops all containers in a container group. Compute resources will be deallocated and billing will stop.
 * x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2021-10-01/examples/ContainerGroupsStop.json
 */
async function containerStop() {
  const subscriptionId = "subid";
  const resourceGroupName = "demo";
  const containerGroupName = "demo1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerInstanceManagementClient(credential, subscriptionId);
  const result = await client.containerGroups.stop(resourceGroupName, containerGroupName);
  console.log(result);
}

containerStop().catch(console.error);
