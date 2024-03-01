const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets details about the specified streaming job.
 *
 * @summary Gets details about the specified streaming job.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/StreamingJob_Get_Expand.json
 */
async function getAStreamingJobAndUseTheExpandODataQueryParameterToExpandInputsOutputsTransformationAndFunctions() {
  const subscriptionId =
    process.env["STREAMANALYTICS_SUBSCRIPTION_ID"] || "56b5e0a9-b645-407d-99b0-c64f86013e3d";
  const expand = "inputs,outputs,transformation,functions";
  const resourceGroupName = process.env["STREAMANALYTICS_RESOURCE_GROUP"] || "sjrg3276";
  const jobName = "sj7804";
  const options = { expand };
  const credential = new DefaultAzureCredential();
  const client = new StreamAnalyticsManagementClient(credential, subscriptionId);
  const result = await client.streamingJobs.get(resourceGroupName, jobName, options);
  console.log(result);
}
