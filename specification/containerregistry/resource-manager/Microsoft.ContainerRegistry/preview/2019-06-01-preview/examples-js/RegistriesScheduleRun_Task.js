const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Schedules a new run based on the request parameters and add it to the run queue.
 *
 * @summary Schedules a new run based on the request parameters and add it to the run queue.
 * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/RegistriesScheduleRun_Task.json
 */
async function registriesScheduleRunTask() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const resourceGroupName = "myResourceGroup";
  const registryName = "myRegistry";
  const runRequest = {
    type: "TaskRunRequest",
    overrideTaskStepProperties: {
      arguments: [
        { name: "mytestargument", isSecret: false, value: "mytestvalue" },
        {
          name: "mysecrettestargument",
          isSecret: true,
          value: "mysecrettestvalue",
        },
      ],
      file: "overriddenDockerfile",
      target: "build",
      updateTriggerToken: "aGVsbG8gd29ybGQ=",
      values: [
        { name: "mytestname", isSecret: false, value: "mytestvalue" },
        { name: "mysecrettestname", isSecret: true, value: "mysecrettestvalue" },
      ],
    },
    taskId: "myTask",
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

registriesScheduleRunTask().catch(console.error);
