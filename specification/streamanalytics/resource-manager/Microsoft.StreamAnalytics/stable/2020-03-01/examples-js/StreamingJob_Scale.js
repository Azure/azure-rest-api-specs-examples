const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Scales a streaming job when the job is running.
 *
 * @summary Scales a streaming job when the job is running.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/StreamingJob_Scale.json
 */
async function scaleAStreamingJob() {
  const subscriptionId = "56b5e0a9-b645-407d-99b0-c64f86013e3d";
  const resourceGroupName = "sjrg6936";
  const jobName = "sj59";
  const scaleJobParameters = {
    streamingUnits: 36,
  };
  const options = { scaleJobParameters };
  const credential = new DefaultAzureCredential();
  const client = new StreamAnalyticsManagementClient(credential, subscriptionId);
  const result = await client.streamingJobs.beginScaleAndWait(resourceGroupName, jobName, options);
  console.log(result);
}

scaleAStreamingJob().catch(console.error);
