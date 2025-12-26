const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Schedules a new run based on the request parameters and add it to the run queue.
 *
 * @summary Schedules a new run based on the request parameters and add it to the run queue.
 * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/RegistryTasks/preview/2025-03-01-preview/examples/RegistriesScheduleRun_EncodedTaskRun.json
 */
async function registriesScheduleRunEncodedTaskRun() {
  const subscriptionId =
    process.env["CONTAINERREGISTRY_SUBSCRIPTION_ID"] || "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const resourceGroupName = process.env["CONTAINERREGISTRY_RESOURCE_GROUP"] || "myResourceGroup";
  const registryName = "myRegistry";
  const runRequest = {
    type: "EncodedTaskRunRequest",
    agentConfiguration: { cpu: 2 },
    encodedTaskContent:
      "c3RlcHM6Cnt7IGlmIFZhbHVlcy5lbnZpcm9ubWVudCA9PSAncHJvZCcgfX0KICAtIHJ1bjogcHJvZCBzZXR1cAp7eyBlbHNlIGlmIFZhbHVlcy5lbnZpcm9ubWVudCA9PSAnc3RhZ2luZycgfX0KICAtIHJ1bjogc3RhZ2luZyBzZXR1cAp7eyBlbHNlIH19CiAgLSBydW46IGRlZmF1bHQgc2V0dXAKe3sgZW5kIH19CgogIC0gcnVuOiBidWlsZCAtdCBGYW5jeVRoaW5nOnt7LlZhbHVlcy5lbnZpcm9ubWVudH19LXt7LlZhbHVlcy52ZXJzaW9ufX0gLgoKcHVzaDogWydGYW5jeVRoaW5nOnt7LlZhbHVlcy5lbnZpcm9ubWVudH19LXt7LlZhbHVlcy52ZXJzaW9ufX0nXQ==",
    encodedValuesContent: "ZW52aXJvbm1lbnQ6IHByb2QKdmVyc2lvbjogMQ==",
    platform: { os: "Linux" },
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
