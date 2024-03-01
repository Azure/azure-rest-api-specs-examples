const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the streaming jobs in the given subscription.
 *
 * @summary Lists all of the streaming jobs in the given subscription.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/StreamingJob_List_BySubscription_Expand.json
 */
async function listAllStreamingJobsInASubscriptionAndUseTheExpandODataQueryParameterToExpandInputsOutputsTransformationAndFunctions() {
  const subscriptionId =
    process.env["STREAMANALYTICS_SUBSCRIPTION_ID"] || "56b5e0a9-b645-407d-99b0-c64f86013e3d";
  const expand = "inputs,outputs,transformation,functions";
  const options = { expand };
  const credential = new DefaultAzureCredential();
  const client = new StreamAnalyticsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.streamingJobs.list(options)) {
    resArray.push(item);
  }
  console.log(resArray);
}
