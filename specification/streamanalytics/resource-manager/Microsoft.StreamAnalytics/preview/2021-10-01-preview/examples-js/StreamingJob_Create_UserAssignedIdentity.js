const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a streaming job or replaces an already existing streaming job.
 *
 * @summary Creates a streaming job or replaces an already existing streaming job.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/StreamingJob_Create_UserAssignedIdentity.json
 */
async function createAStreamingJobShellAStreamingJobWithNoInputsOutputsTransformationOrFunctionsWithUserAssignedIdentity() {
  const subscriptionId =
    process.env["STREAMANALYTICS_SUBSCRIPTION_ID"] || "56b5e0a9-b645-407d-99b0-c64f86013e3d";
  const resourceGroupName = process.env["STREAMANALYTICS_RESOURCE_GROUP"] || "sjrg";
  const jobName = "sjName";
  const streamingJob = {
    compatibilityLevel: "1.0",
    dataLocale: "en-US",
    eventsLateArrivalMaxDelayInSeconds: 16,
    eventsOutOfOrderMaxDelayInSeconds: 5,
    eventsOutOfOrderPolicy: "Drop",
    functions: [],
    identity: {
      type: "UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/fa68082f8ff74a2595c7Ce9da541242f/resourceGroups/akvenkat/providers/MicrosoftManagedIdentity/userAssignedIdentities/sdkIdentity":
          {},
      },
    },
    inputs: [],
    location: "West US",
    outputErrorPolicy: "Drop",
    outputs: [],
    skuPropertiesSku: { name: "Standard" },
    tags: { key1: "value1", key3: "value3", randomKey: "randomValue" },
  };
  const credential = new DefaultAzureCredential();
  const client = new StreamAnalyticsManagementClient(credential, subscriptionId);
  const result = await client.streamingJobs.beginCreateOrReplaceAndWait(
    resourceGroupName,
    jobName,
    streamingJob,
  );
  console.log(result);
}
