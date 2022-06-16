const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an existing input under an existing streaming job. This can be used to partially update (ie. update one or two properties) an input without affecting the rest the job or input definition.
 *
 * @summary Updates an existing input under an existing streaming job. This can be used to partially update (ie. update one or two properties) an input without affecting the rest the job or input definition.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Input_Update_Stream_EventHub.json
 */
async function updateAStreamEventHubInput() {
  const subscriptionId = "56b5e0a9-b645-407d-99b0-c64f86013e3d";
  const resourceGroupName = "sjrg3139";
  const jobName = "sj197";
  const inputName = "input7425";
  const input = {
    properties: {
      type: "Stream",
      datasource: {
        type: "Microsoft.ServiceBus/EventHub",
        consumerGroupName: "differentConsumerGroupName",
      },
      serialization: { type: "Avro" },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new StreamAnalyticsManagementClient(credential, subscriptionId);
  const result = await client.inputs.update(resourceGroupName, jobName, inputName, input);
  console.log(result);
}

updateAStreamEventHubInput().catch(console.error);
