const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets details about the specified output.
 *
 * @summary Gets details about the specified output.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Output_Get_Blob_CSV.json
 */
async function getABlobOutputWithCsvSerialization() {
  const subscriptionId = "56b5e0a9-b645-407d-99b0-c64f86013e3d";
  const resourceGroupName = "sjrg5023";
  const jobName = "sj900";
  const outputName = "output1623";
  const credential = new DefaultAzureCredential();
  const client = new StreamAnalyticsManagementClient(credential, subscriptionId);
  const result = await client.outputs.get(resourceGroupName, jobName, outputName);
  console.log(result);
}

getABlobOutputWithCsvSerialization().catch(console.error);
