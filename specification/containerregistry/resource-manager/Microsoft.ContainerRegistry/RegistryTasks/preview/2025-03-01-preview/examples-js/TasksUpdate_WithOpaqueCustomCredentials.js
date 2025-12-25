const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Updates a task with the specified parameters.
 *
 * @summary Updates a task with the specified parameters.
 * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/RegistryTasks/preview/2025-03-01-preview/examples/TasksUpdate_WithOpaqueCustomCredentials.json
 */
async function tasksUpdateWithOpaqueCustomCredentials() {
  const subscriptionId =
    process.env["CONTAINERREGISTRY_SUBSCRIPTION_ID"] || "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const resourceGroupName = process.env["CONTAINERREGISTRY_RESOURCE_GROUP"] || "myResourceGroup";
  const registryName = "myRegistry";
  const taskName = "myTask";
  const taskUpdateParameters = {
    agentConfiguration: { cpu: 3 },
    credentials: {
      customRegistries: {
        myregistryAzurecrIo: {
          password: { type: "Opaque", value: "***" },
          userName: { type: "Opaque", value: "username" },
        },
      },
    },
    logTemplate: undefined,
    status: "Enabled",
    step: {
      type: "Docker",
      dockerFilePath: "src/DockerFile",
      imageNames: ["azurerest:testtag1"],
    },
    tags: { testkey: "value" },
    trigger: {
      sourceTriggers: [
        {
          name: "mySourceTrigger",
          sourceRepository: {
            sourceControlAuthProperties: { token: "xxxxx", tokenType: "PAT" },
          },
          sourceTriggerEvents: ["commit"],
        },
      ],
    },
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
