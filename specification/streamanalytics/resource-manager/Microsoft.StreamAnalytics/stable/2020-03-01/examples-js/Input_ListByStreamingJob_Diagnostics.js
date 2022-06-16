const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the inputs under the specified streaming job.
 *
 * @summary Lists all of the inputs under the specified streaming job.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Input_ListByStreamingJob_Diagnostics.json
 */
async function listAllInputsInAStreamingJobAndIncludeDiagnosticInformationUsingTheSelectODataQueryParameter() {
  const subscriptionId = "56b5e0a9-b645-407d-99b0-c64f86013e3d";
  const select = "*";
  const resourceGroupName = "sjrg3276";
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

listAllInputsInAStreamingJobAndIncludeDiagnosticInformationUsingTheSelectODataQueryParameter().catch(
  console.error
);
