const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the functions under the specified streaming job.
 *
 * @summary Lists all of the functions under the specified streaming job.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Function_ListByStreamingJob.json
 */
async function listAllFunctionsInAStreamingJob() {
  const subscriptionId = "56b5e0a9-b645-407d-99b0-c64f86013e3d";
  const resourceGroupName = "sjrg1637";
  const jobName = "sj8653";
  const credential = new DefaultAzureCredential();
  const client = new StreamAnalyticsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.functions.listByStreamingJob(resourceGroupName, jobName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAllFunctionsInAStreamingJob().catch(console.error);
