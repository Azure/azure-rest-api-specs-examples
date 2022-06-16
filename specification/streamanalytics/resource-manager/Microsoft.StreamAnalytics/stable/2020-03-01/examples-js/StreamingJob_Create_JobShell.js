const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a streaming job or replaces an already existing streaming job.
 *
 * @summary Creates a streaming job or replaces an already existing streaming job.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/StreamingJob_Create_JobShell.json
 */
async function createAStreamingJobShellAStreamingJobWithNoInputsOutputsTransformationOrFunctions() {
  const subscriptionId = "56b5e0a9-b645-407d-99b0-c64f86013e3d";
  const resourceGroupName = "sjrg6936";
  const jobName = "sj59";
  const streamingJob = {
    compatibilityLevel: "1.0",
    dataLocale: "en-US",
    eventsLateArrivalMaxDelayInSeconds: 16,
    eventsOutOfOrderMaxDelayInSeconds: 5,
    eventsOutOfOrderPolicy: "Drop",
    functions: [],
    inputs: [],
    location: "West US",
    outputErrorPolicy: "Drop",
    outputs: [],
    sku: { name: "Standard" },
    tags: { key1: "value1", key3: "value3", randomKey: "randomValue" },
  };
  const credential = new DefaultAzureCredential();
  const client = new StreamAnalyticsManagementClient(credential, subscriptionId);
  const result = await client.streamingJobs.beginCreateOrReplaceAndWait(
    resourceGroupName,
    jobName,
    streamingJob
  );
  console.log(result);
}

createAStreamingJobShellAStreamingJobWithNoInputsOutputsTransformationOrFunctions().catch(
  console.error
);
