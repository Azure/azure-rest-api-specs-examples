const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the inputs under the specified streaming job.
 *
 * @summary Lists all of the inputs under the specified streaming job.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Input_ListByStreamingJob_Diagnostics.json
 */
async function listAllInputsInAStreamingJobAndIncludeDiagnosticInformationUsingTheSelectODataQueryParameter() {
  const subscriptionId =
    process.env["STREAMANALYTICS_SUBSCRIPTION_ID"] || "56b5e0a9-b645-407d-99b0-c64f86013e3d";
  const select = "*";
  const resourceGroupName = process.env["STREAMANALYTICS_RESOURCE_GROUP"] || "sjrg3276";
  const jobName = "sj7804";
  const options = { select };
  const credential = new DefaultAzureCredential();
  const client = new StreamAnalyticsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.inputs.listByStreamingJob(resourceGroupName, jobName, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}
