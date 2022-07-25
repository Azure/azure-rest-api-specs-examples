const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a task with the specified parameters.
 *
 * @summary Updates a task with the specified parameters.
 * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/TasksUpdate_QuickTask.json
 */
async function tasksUpdateQuickTask() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const resourceGroupName = "myResourceGroup";
  const registryName = "myRegistry";
  const taskName = "quicktask";
  const taskUpdateParameters = {
    logTemplate: "acr/tasks:{{.Run.OS}}",
    status: "Enabled",
    tags: { testkey: "value" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerRegistryManagementClient(credential, subscriptionId);
  const result = await client.tasks.beginUpdateAndWait(
    resourceGroupName,
    registryName,
    taskName,
    taskUpdateParameters
  );
  console.log(result);
}

tasksUpdateQuickTask().catch(console.error);
