const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets details about the specified input.
 *
 * @summary Gets details about the specified input.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Input_Get_Reference_Blob_CSV.json
 */
async function getAReferenceBlobInputWithCsvSerialization() {
  const subscriptionId = "56b5e0a9-b645-407d-99b0-c64f86013e3d";
  const resourceGroupName = "sjrg8440";
  const jobName = "sj9597";
  const inputName = "input7225";
  const credential = new DefaultAzureCredential();
  const client = new StreamAnalyticsManagementClient(credential, subscriptionId);
  const result = await client.inputs.get(resourceGroupName, jobName, inputName);
  console.log(result);
}

getAReferenceBlobInputWithCsvSerialization().catch(console.error);
