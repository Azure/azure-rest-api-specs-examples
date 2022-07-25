const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Schedules a new run based on the request parameters and add it to the run queue.
 *
 * @summary Schedules a new run based on the request parameters and add it to the run queue.
 * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/RegistriesScheduleRun_FileTask_WithCustomCredentials.json
 */
async function registriesScheduleRunTaskWithCustomCredentials() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const resourceGroupName = "myResourceGroup";
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
  const result = await client.registries.beginScheduleRunAndWait(
    resourceGroupName,
    registryName,
    runRequest
  );
  console.log(result);
}

registriesScheduleRunTaskWithCustomCredentials().catch(console.error);
