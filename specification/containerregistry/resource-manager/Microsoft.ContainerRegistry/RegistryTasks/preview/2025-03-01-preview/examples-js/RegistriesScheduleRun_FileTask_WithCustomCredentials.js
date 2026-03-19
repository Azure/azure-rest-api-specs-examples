const { ContainerRegistryTasksManagementClient } = require("@azure/arm-containerregistrytasks");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to schedules a new run based on the request parameters and add it to the run queue.
 *
 * @summary schedules a new run based on the request parameters and add it to the run queue.
 * x-ms-original-file: 2025-03-01-preview/RegistriesScheduleRun_FileTask_WithCustomCredentials.json
 */
async function registriesScheduleRunTaskWithCustomCredentials() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const client = new ContainerRegistryTasksManagementClient(credential, subscriptionId);
  const result = await client.registries.scheduleRun("myResourceGroup", "myRegistry", {
    type: "FileTaskRunRequest",
    credentials: {
      customRegistries: {
        "myregistry.azurecr.io": {
          password: { type: "Opaque", value: "***" },
          userName: { type: "Opaque", value: "reg1" },
        },
      },
      sourceRegistry: { loginMode: "Default" },
    },
    platform: { os: "Linux" },
    taskFilePath: "acb.yaml",
    values: [
      { name: "mytestargument", isSecret: false, value: "mytestvalue" },
      { name: "mysecrettestargument", isSecret: true, value: "mysecrettestvalue" },
    ],
  });
  console.log(result);
}
