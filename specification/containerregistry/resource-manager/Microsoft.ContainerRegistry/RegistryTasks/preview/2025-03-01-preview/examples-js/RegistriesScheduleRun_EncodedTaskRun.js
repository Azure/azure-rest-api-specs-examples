const { ContainerRegistryTasksManagementClient } = require("@azure/arm-containerregistrytasks");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to schedules a new run based on the request parameters and add it to the run queue.
 *
 * @summary schedules a new run based on the request parameters and add it to the run queue.
 * x-ms-original-file: 2025-03-01-preview/RegistriesScheduleRun_EncodedTaskRun.json
 */
async function registriesScheduleRunEncodedTaskRun() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const client = new ContainerRegistryTasksManagementClient(credential, subscriptionId);
  const result = await client.registries.scheduleRun("myResourceGroup", "myRegistry", {
    type: "EncodedTaskRunRequest",
    agentConfiguration: { cpu: 2 },
    encodedTaskContent:
      "c3RlcHM6Cnt7IGlmIFZhbHVlcy5lbnZpcm9ubWVudCA9PSAncHJvZCcgfX0KICAtIHJ1bjogcHJvZCBzZXR1cAp7eyBlbHNlIGlmIFZhbHVlcy5lbnZpcm9ubWVudCA9PSAnc3RhZ2luZycgfX0KICAtIHJ1bjogc3RhZ2luZyBzZXR1cAp7eyBlbHNlIH19CiAgLSBydW46IGRlZmF1bHQgc2V0dXAKe3sgZW5kIH19CgogIC0gcnVuOiBidWlsZCAtdCBGYW5jeVRoaW5nOnt7LlZhbHVlcy5lbnZpcm9ubWVudH19LXt7LlZhbHVlcy52ZXJzaW9ufX0gLgoKcHVzaDogWydGYW5jeVRoaW5nOnt7LlZhbHVlcy5lbnZpcm9ubWVudH19LXt7LlZhbHVlcy52ZXJzaW9ufX0nXQ==",
    encodedValuesContent: "ZW52aXJvbm1lbnQ6IHByb2QKdmVyc2lvbjogMQ==",
    platform: { os: "Linux" },
    values: [
      { name: "mytestargument", isSecret: false, value: "mytestvalue" },
      { name: "mysecrettestargument", isSecret: true, value: "mysecrettestvalue" },
    ],
  });
  console.log(result);
}
