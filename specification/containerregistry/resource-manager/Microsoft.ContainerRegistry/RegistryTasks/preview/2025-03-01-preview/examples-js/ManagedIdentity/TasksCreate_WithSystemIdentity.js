const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates a task for a container registry with the specified parameters.
 *
 * @summary Creates a task for a container registry with the specified parameters.
 * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/RegistryTasks/preview/2025-03-01-preview/examples/ManagedIdentity/TasksCreate_WithSystemIdentity.json
 */
async function tasksCreateWithUserIdentitiesWithSystemIdentity() {
  const subscriptionId =
    process.env["CONTAINERREGISTRY_SUBSCRIPTION_ID"] || "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const resourceGroupName = process.env["CONTAINERREGISTRY_RESOURCE_GROUP"] || "myResourceGroup";
  const registryName = "myRegistry";
  const taskName = "mytTask";
  const taskCreateParameters = {
    agentConfiguration: { cpu: 2 },
    identity: { type: "SystemAssigned" },
    isSystemTask: false,
    location: "eastus",
    logTemplate: undefined,
    platform: { architecture: "amd64", os: "Linux" },
    status: "Enabled",
    step: {
      type: "Docker",
      arguments: [
        { name: "mytestargument", isSecret: false, value: "mytestvalue" },
        {
          name: "mysecrettestargument",
          isSecret: true,
          value: "mysecrettestvalue",
        },
      ],
      contextPath: "src",
      dockerFilePath: "src/DockerFile",
      imageNames: ["azurerest:testtag"],
      isPushEnabled: true,
      noCache: false,
    },
    tags: { testkey: "value" },
    trigger: {
      baseImageTrigger: {
        name: "myBaseImageTrigger",
        baseImageTriggerType: "Runtime",
      },
      sourceTriggers: [
        {
          name: "mySourceTrigger",
          sourceRepository: {
            branch: "master",
            repositoryUrl: "https://github.com/Azure/azure-rest-api-specs",
            sourceControlAuthProperties: { token: "xxxxx", tokenType: "PAT" },
            sourceControlType: "Github",
          },
          sourceTriggerEvents: ["commit"],
        },
      ],
      timerTriggers: [{ name: "myTimerTrigger", schedule: "30 9 * * 1-5" }],
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerRegistryManagementClient(credential, subscriptionId);
  const result = await client.tasks.create(
    resourceGroupName,
    registryName,
    taskName,
    taskCreateParameters,
  );
  console.log(result);
}
