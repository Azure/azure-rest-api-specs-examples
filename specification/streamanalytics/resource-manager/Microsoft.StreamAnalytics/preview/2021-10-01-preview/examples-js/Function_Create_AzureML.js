const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a function or replaces an already existing function under an existing streaming job.
 *
 * @summary Creates a function or replaces an already existing function under an existing streaming job.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Function_Create_AzureML.json
 */
async function createAnAzureMlFunction() {
  const subscriptionId =
    process.env["STREAMANALYTICS_SUBSCRIPTION_ID"] || "56b5e0a9-b645-407d-99b0-c64f86013e3d";
  const resourceGroupName = process.env["STREAMANALYTICS_RESOURCE_GROUP"] || "sjrg7";
  const jobName = "sj9093";
  const functionName = "function588";
  const functionParam = {
    properties: {
      type: "Scalar",
      binding: {
        type: "Microsoft.MachineLearning/WebService",
        apiKey: "someApiKey==",
        batchSize: 1000,
        endpoint: "someAzureMLEndpointURL",
        inputs: {
          name: "input1",
          columnNames: [{ name: "tweet", dataType: "string", mapTo: 0 }],
        },
        outputs: [{ name: "Sentiment", dataType: "string" }],
      },
      inputs: [{ dataType: "nvarchar(max)" }],
      output: { dataType: "nvarchar(max)" },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new StreamAnalyticsManagementClient(credential, subscriptionId);
  const result = await client.functions.createOrReplace(
    resourceGroupName,
    jobName,
    functionName,
    functionParam,
  );
  console.log(result);
}
