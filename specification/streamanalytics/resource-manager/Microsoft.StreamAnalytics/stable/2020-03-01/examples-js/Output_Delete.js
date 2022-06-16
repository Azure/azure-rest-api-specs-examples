const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes an output from the streaming job.
 *
 * @summary Deletes an output from the streaming job.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Output_Delete.json
 */
async function deleteAnOutput() {
  const subscriptionId = "56b5e0a9-b645-407d-99b0-c64f86013e3d";
  const resourceGroupName = "sjrg2157";
  const jobName = "sj6458";
  const outputName = "output1755";
  const credential = new DefaultAzureCredential();
  const client = new StreamAnalyticsManagementClient(credential, subscriptionId);
  const result = await client.outputs.delete(resourceGroupName, jobName, outputName);
  console.log(result);
}

deleteAnOutput().catch(console.error);
