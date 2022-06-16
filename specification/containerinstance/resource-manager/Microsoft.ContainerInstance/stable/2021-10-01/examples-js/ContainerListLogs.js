const { ContainerInstanceManagementClient } = require("@azure/arm-containerinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the logs for a specified container instance in a specified resource group and container group.
 *
 * @summary Get the logs for a specified container instance in a specified resource group and container group.
 * x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2021-10-01/examples/ContainerListLogs.json
 */
async function containerListLogs() {
  const subscriptionId = "subid";
  const resourceGroupName = "demo";
  const containerGroupName = "demo1";
  const containerName = "container1";
  const tail = 10;
  const options = { tail };
  const credential = new DefaultAzureCredential();
  const client = new ContainerInstanceManagementClient(credential, subscriptionId);
  const result = await client.containers.listLogs(
    resourceGroupName,
    containerGroupName,
    containerName,
    options
  );
  console.log(result);
}

containerListLogs().catch(console.error);
