const { ContainerRegistryTasksManagementClient } = require("@azure/arm-containerregistrytasks");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a task for a container registry with the specified parameters.
 *
 * @summary creates a task for a container registry with the specified parameters.
 * x-ms-original-file: 2025-03-01-preview/TasksCreate_QuickTask.json
 */
async function tasksCreateQuickTask() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const client = new ContainerRegistryTasksManagementClient(credential, subscriptionId);
  const result = await client.tasks.create("myResourceGroup", "myRegistry", "quicktask", {
    location: "eastus",
    properties: { isSystemTask: true, logTemplate: "acr/tasks:{{.Run.OS}}", status: "Enabled" },
    tags: { testkey: "value" },
  });
  console.log(result);
}
