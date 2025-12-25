const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Schedules a new run based on the request parameters and add it to the run queue.
 *
 * @summary Schedules a new run based on the request parameters and add it to the run queue.
 * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/RegistryTasks/preview/2025-03-01-preview/examples/RegistriesScheduleRun_FileTask_WithCustomCredentials.json
 */
async function registriesScheduleRunTaskWithCustomCredentials() {
  const subscriptionId =
    process.env["CONTAINERREGISTRY_SUBSCRIPTION_ID"] || "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const resourceGroupName = process.env["CONTAINERREGISTRY_RESOURCE_GROUP"] || "myResourceGroup";
  const registryName = "myRegistry";
  const runRequest = {
    type: "FileTaskRunRequest",
    credentials: {
      customRegistries: {
        myregistryAzurecrIo: {
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
      {
        name: "mysecrettestargument",
        isSecret: true,
        value: "mysecrettestvalue",
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerRegistryManagementClient(credential, subscriptionId);
  const result = await client.registries.scheduleRun(resourceGroupName, registryName, runRequest);
  console.log(result);
}
