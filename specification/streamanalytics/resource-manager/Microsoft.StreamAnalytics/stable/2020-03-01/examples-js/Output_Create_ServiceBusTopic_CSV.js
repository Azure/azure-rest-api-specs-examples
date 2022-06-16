const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates an output or replaces an already existing output under an existing streaming job.
 *
 * @summary Creates an output or replaces an already existing output under an existing streaming job.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Output_Create_ServiceBusTopic_CSV.json
 */
async function createAServiceBusTopicOutputWithCsvSerialization() {
  const subscriptionId = "56b5e0a9-b645-407d-99b0-c64f86013e3d";
  const resourceGroupName = "sjrg6450";
  const jobName = "sj7094";
  const outputName = "output7886";
  const output = {
    datasource: {
      type: "Microsoft.ServiceBus/Topic",
      propertyColumns: ["column1", "column2"],
      serviceBusNamespace: "sdktest",
      sharedAccessPolicyKey: "sharedAccessPolicyKey=",
      sharedAccessPolicyName: "RootManageSharedAccessKey",
      topicName: "sdktopic",
    },
    serialization: { type: "Csv", encoding: "UTF8", fieldDelimiter: "," },
  };
  const credential = new DefaultAzureCredential();
  const client = new StreamAnalyticsManagementClient(credential, subscriptionId);
  const result = await client.outputs.createOrReplace(
    resourceGroupName,
    jobName,
    outputName,
    output
  );
  console.log(result);
}

createAServiceBusTopicOutputWithCsvSerialization().catch(console.error);
