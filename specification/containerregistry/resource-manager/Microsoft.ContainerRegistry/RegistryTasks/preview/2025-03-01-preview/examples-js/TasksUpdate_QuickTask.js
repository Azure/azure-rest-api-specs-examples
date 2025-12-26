const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Updates a task with the specified parameters.
 *
 * @summary Updates a task with the specified parameters.
 * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/RegistryTasks/preview/2025-03-01-preview/examples/TasksUpdate_QuickTask.json
 */
async function tasksUpdateQuickTask() {
  const subscriptionId =
    process.env["CONTAINERREGISTRY_SUBSCRIPTION_ID"] || "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const resourceGroupName = process.env["CONTAINERREGISTRY_RESOURCE_GROUP"] || "myResourceGroup";
  const registryName = "myRegistry";
  const taskName = "quicktask";
  const taskUpdateParameters = {
    logTemplate: "acr/tasks:{{.Run.OS}}",
    status: "Enabled",
    tags: { testkey: "value" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerRegistryManagementClient(credential, subscriptionId);
  const result = await client.tasks.update(
    resourceGroupName,
    registryName,
    taskName,
    taskUpdateParameters,
  );
  console.log(result);
}
