const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a task for a container registry with the specified parameters.
 *
 * @summary Creates a task for a container registry with the specified parameters.
 * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/TasksCreate_QuickTask.json
 */
async function tasksCreateQuickTask() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const resourceGroupName = "myResourceGroup";
  const registryName = "myRegistry";
  const taskName = "quicktask";
  const taskCreateParameters = {
    identity: {},
    isSystemTask: true,
    location: "eastus",
    logTemplate: "acr/tasks:{{.Run.OS}}",
    status: "Enabled",
    tags: { testkey: "value" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerRegistryManagementClient(credential, subscriptionId);
  const result = await client.tasks.beginCreateAndWait(
    resourceGroupName,
    registryName,
    taskName,
    taskCreateParameters
  );
  console.log(result);
}

tasksCreateQuickTask().catch(console.error);
