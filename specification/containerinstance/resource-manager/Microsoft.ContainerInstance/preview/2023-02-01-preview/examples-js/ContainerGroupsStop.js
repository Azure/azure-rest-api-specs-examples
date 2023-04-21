const { ContainerInstanceManagementClient } = require("@azure/arm-containerinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Stops all containers in a container group. Compute resources will be deallocated and billing will stop.
 *
 * @summary Stops all containers in a container group. Compute resources will be deallocated and billing will stop.
 * x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/preview/2023-02-01-preview/examples/ContainerGroupsStop.json
 */
async function containerStop() {
  const subscriptionId = process.env["CONTAINERINSTANCE_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["CONTAINERINSTANCE_RESOURCE_GROUP"] || "demo";
  const containerGroupName = "demo1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerInstanceManagementClient(credential, subscriptionId);
  const result = await client.containerGroups.stop(resourceGroupName, containerGroupName);
  console.log(result);
}
