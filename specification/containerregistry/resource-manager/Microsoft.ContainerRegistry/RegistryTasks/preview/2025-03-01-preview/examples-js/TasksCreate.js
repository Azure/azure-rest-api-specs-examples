const { ContainerRegistryTasksManagementClient } = require("@azure/arm-containerregistrytasks");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a task for a container registry with the specified parameters.
 *
 * @summary creates a task for a container registry with the specified parameters.
 * x-ms-original-file: 2025-03-01-preview/TasksCreate.json
 */
async function tasksCreate() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const client = new ContainerRegistryTasksManagementClient(credential, subscriptionId);
  const result = await client.tasks.create("myResourceGroup", "myRegistry", "mytTask", {
    identity: { type: "SystemAssigned" },
    location: "eastus",
    properties: {
      agentConfiguration: { cpu: 2 },
      isSystemTask: false,
      logTemplate: "acr/tasks:{{.Run.OS}}",
      platform: { architecture: "amd64", os: "Linux" },
      status: "Enabled",
      step: {
        type: "Docker",
        arguments: [
          { name: "mytestargument", isSecret: false, value: "mytestvalue" },
          { name: "mysecrettestargument", isSecret: true, value: "mysecrettestvalue" },
        ],
        contextPath: "src",
        dockerFilePath: "src/DockerFile",
        imageNames: ["azurerest:testtag"],
        isPushEnabled: true,
        noCache: false,
      },
      trigger: {
        baseImageTrigger: {
          name: "myBaseImageTrigger",
          baseImageTriggerType: "Runtime",
          updateTriggerEndpoint: "https://user:pass@mycicd.webhook.com?token=foo",
          updateTriggerPayloadType: "Token",
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
    },
    tags: { testkey: "value" },
  });
  console.log(result);
}
