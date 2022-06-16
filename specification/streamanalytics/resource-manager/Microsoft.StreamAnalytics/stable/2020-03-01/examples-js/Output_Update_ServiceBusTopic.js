const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an existing output under an existing streaming job. This can be used to partially update (ie. update one or two properties) an output without affecting the rest the job or output definition.
 *
 * @summary Updates an existing output under an existing streaming job. This can be used to partially update (ie. update one or two properties) an output without affecting the rest the job or output definition.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Output_Update_ServiceBusTopic.json
 */
async function updateAServiceBusTopicOutputWithCsvSerialization() {
  const subscriptionId = "56b5e0a9-b645-407d-99b0-c64f86013e3d";
  const resourceGroupName = "sjrg6450";
  const jobName = "sj7094";
  const outputName = "output7886";
  const output = {
    datasource: {
      type: "Microsoft.ServiceBus/Topic",
      topicName: "differentTopicName",
    },
    serialization: { type: "Csv", encoding: "UTF8", fieldDelimiter: "|" },
  };
  const credential = new DefaultAzureCredential();
  const client = new StreamAnalyticsManagementClient(credential, subscriptionId);
  const result = await client.outputs.update(resourceGroupName, jobName, outputName, output);
  console.log(result);
}

updateAServiceBusTopicOutputWithCsvSerialization().catch(console.error);
