const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Starts a streaming job. Once a job is started it will start processing input events and produce output.
 *
 * @summary Starts a streaming job. Once a job is started it will start processing input events and produce output.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/StreamingJob_Start_LastOutputEventTime.json
 */
async function startAStreamingJobWithLastOutputEventTimeOutputStartMode() {
  const subscriptionId = "56b5e0a9-b645-407d-99b0-c64f86013e3d";
  const resourceGroupName = "sjrg6936";
  const jobName = "sj59";
  const startJobParameters = {
    outputStartMode: "LastOutputEventTime",
  };
  const options = { startJobParameters };
  const credential = new DefaultAzureCredential();
  const client = new StreamAnalyticsManagementClient(credential, subscriptionId);
  const result = await client.streamingJobs.beginStartAndWait(resourceGroupName, jobName, options);
  console.log(result);
}

startAStreamingJobWithLastOutputEventTimeOutputStartMode().catch(console.error);
