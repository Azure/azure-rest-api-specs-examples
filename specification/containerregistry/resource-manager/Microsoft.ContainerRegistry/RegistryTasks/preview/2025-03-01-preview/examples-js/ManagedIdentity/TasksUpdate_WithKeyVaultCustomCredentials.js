const { ContainerRegistryTasksManagementClient } = require("@azure/arm-containerregistrytasks");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates a task with the specified parameters.
 *
 * @summary updates a task with the specified parameters.
 * x-ms-original-file: 2025-03-01-preview/ManagedIdentity/TasksUpdate_WithKeyVaultCustomCredentials.json
 */
async function tasksUpdateWithKeyVaultCustomCredentials() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const client = new ContainerRegistryTasksManagementClient(credential, subscriptionId);
  const result = await client.tasks.update("myResourceGroup", "myRegistry", "myTask", {
    agentConfiguration: { cpu: 3 },
    credentials: {
      customRegistries: {
        "myregistry.azurecr.io": {
          identity: "[system]",
          password: {
            type: "Vaultsecret",
            value: "https://myacbvault.vault.azure.net/secrets/password",
          },
          userName: {
            type: "Vaultsecret",
            value: "https://myacbvault.vault.azure.net/secrets/username",
          },
        },
      },
    },
    status: "Enabled",
    step: { type: "Docker", dockerFilePath: "src/DockerFile", imageNames: ["azurerest:testtag1"] },
    trigger: {
      sourceTriggers: [
        {
          name: "mySourceTrigger",
          sourceRepository: { sourceControlAuthProperties: { token: "xxxxx", tokenType: "PAT" } },
          sourceTriggerEvents: ["commit"],
        },
      ],
    },
    tags: { testkey: "value" },
  });
  console.log(result);
}
