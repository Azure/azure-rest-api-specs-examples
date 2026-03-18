const { ContainerRegistryTasksManagementClient } = require("@azure/arm-containerregistrytasks");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to schedules a new run based on the request parameters and add it to the run queue.
 *
 * @summary schedules a new run based on the request parameters and add it to the run queue.
 * x-ms-original-file: 2025-03-01-preview/RegistriesScheduleRun_FileTaskRun.json
 */
async function registriesScheduleRunFileTaskRun() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const client = new ContainerRegistryTasksManagementClient(credential, subscriptionId);
  const result = await client.registries.scheduleRun("myResourceGroup", "myRegistry", {
    type: "FileTaskRunRequest",
    agentConfiguration: { cpu: 2 },
    platform: { os: "Linux" },
    sourceLocation:
      "https://myaccount.blob.core.windows.net/sascontainer/source.zip?sv=2015-04-05&st=2015-04-29T22%3A18%3A26Z&se=2015-04-30T02%3A23%3A26Z&sr=b&sp=rw&sip=168.1.5.60-168.1.5.70&spr=https&sig=Z%2FRHIX5Xcg0Mq2rqI3OlWTjEg2tYkboXr1P9ZUXDtkk%3D",
    taskFilePath: "acb.yaml",
    values: [
      { name: "mytestargument", isSecret: false, value: "mytestvalue" },
      { name: "mysecrettestargument", isSecret: true, value: "mysecrettestvalue" },
    ],
    valuesFilePath: "prod-values.yaml",
  });
  console.log(result);
}
