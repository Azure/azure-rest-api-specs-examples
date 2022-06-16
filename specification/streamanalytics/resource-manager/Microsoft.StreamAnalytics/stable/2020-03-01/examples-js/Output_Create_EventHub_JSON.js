const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates an output or replaces an already existing output under an existing streaming job.
 *
 * @summary Creates an output or replaces an already existing output under an existing streaming job.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Output_Create_EventHub_JSON.json
 */
async function createAnEventHubOutputWithJsonSerialization() {
  const subscriptionId = "56b5e0a9-b645-407d-99b0-c64f86013e3d";
  const resourceGroupName = "sjrg6912";
  const jobName = "sj3310";
  const outputName = "output5195";
  const output = {
    datasource: {
      type: "Microsoft.ServiceBus/EventHub",
      eventHubName: "sdkeventhub",
      partitionKey: "partitionKey",
      serviceBusNamespace: "sdktest",
      sharedAccessPolicyKey: "sharedAccessPolicyKey=",
      sharedAccessPolicyName: "RootManageSharedAccessKey",
    },
    serialization: { type: "Json", format: "Array", encoding: "UTF8" },
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

createAnEventHubOutputWithJsonSerialization().catch(console.error);
