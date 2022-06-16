const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates an input or replaces an already existing input under an existing streaming job.
 *
 * @summary Creates an input or replaces an already existing input under an existing streaming job.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Input_Create_Stream_EventHub_JSON.json
 */
async function createAStreamEventHubInputWithJsonSerialization() {
  const subscriptionId = "56b5e0a9-b645-407d-99b0-c64f86013e3d";
  const resourceGroupName = "sjrg3139";
  const jobName = "sj197";
  const inputName = "input7425";
  const input = {
    properties: {
      type: "Stream",
      datasource: {
        type: "Microsoft.ServiceBus/EventHub",
        consumerGroupName: "sdkconsumergroup",
        eventHubName: "sdkeventhub",
        serviceBusNamespace: "sdktest",
        sharedAccessPolicyKey: "someSharedAccessPolicyKey==",
        sharedAccessPolicyName: "RootManageSharedAccessKey",
      },
      serialization: { type: "Json", encoding: "UTF8" },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new StreamAnalyticsManagementClient(credential, subscriptionId);
  const result = await client.inputs.createOrReplace(resourceGroupName, jobName, inputName, input);
  console.log(result);
}

createAStreamEventHubInputWithJsonSerialization().catch(console.error);
