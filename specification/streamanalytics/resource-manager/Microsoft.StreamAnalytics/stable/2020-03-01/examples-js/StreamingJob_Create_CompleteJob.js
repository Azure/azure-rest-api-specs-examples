const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a streaming job or replaces an already existing streaming job.
 *
 * @summary Creates a streaming job or replaces an already existing streaming job.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/StreamingJob_Create_CompleteJob.json
 */
async function createACompleteStreamingJobAStreamingJobWithATransformationAtLeast1InputAndAtLeast1Output() {
  const subscriptionId = "56b5e0a9-b645-407d-99b0-c64f86013e3d";
  const resourceGroupName = "sjrg3276";
  const jobName = "sj7804";
  const streamingJob = {
    compatibilityLevel: "1.0",
    dataLocale: "en-US",
    eventsLateArrivalMaxDelayInSeconds: 5,
    eventsOutOfOrderMaxDelayInSeconds: 0,
    eventsOutOfOrderPolicy: "Drop",
    functions: [],
    inputs: [
      {
        name: "inputtest",
        properties: {
          type: "Stream",
          datasource: {
            type: "Microsoft.Storage/Blob",
            container: "containerName",
            pathPattern: "",
            storageAccounts: [{ accountKey: "yourAccountKey==", accountName: "yourAccountName" }],
          },
          serialization: { type: "Json", encoding: "UTF8" },
        },
      },
    ],
    location: "West US",
    outputErrorPolicy: "Drop",
    outputs: [
      {
        name: "outputtest",
        datasource: {
          type: "Microsoft.Sql/Server/Database",
          database: "databaseName",
          password: "userPassword",
          server: "serverName",
          table: "tableName",
          user: "<user>",
        },
      },
    ],
    sku: { name: "Standard" },
    tags: { key1: "value1", key3: "value3", randomKey: "randomValue" },
    transformation: {
      name: "transformationtest",
      query: "Select Id, Name from inputtest",
      streamingUnits: 1,
    },
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

createACompleteStreamingJobAStreamingJobWithATransformationAtLeast1InputAndAtLeast1Output().catch(
  console.error
);
