const { ContainerInstanceManagementClient } = require("@azure/arm-containerinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Restarts all containers in a container group in place. If container image has updates, new image will be downloaded.
 *
 * @summary Restarts all containers in a container group in place. If container image has updates, new image will be downloaded.
 * x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2021-10-01/examples/ContainerGroupsRestart.json
 */
async function containerRestart() {
  const subscriptionId = "subid";
  const resourceGroupName = "demo";
  const containerGroupName = "demo1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerInstanceManagementClient(credential, subscriptionId);
  const result = await client.containerGroups.beginRestartAndWait(
    resourceGroupName,
    containerGroupName
  );
  console.log(result);
}

containerRestart().catch(console.error);
