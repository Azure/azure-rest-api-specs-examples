const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets details about the specified function.
 *
 * @summary Gets details about the specified function.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Function_Get_AzureML.json
 */
async function getAnAzureMlFunction() {
  const subscriptionId = "56b5e0a9-b645-407d-99b0-c64f86013e3d";
  const resourceGroupName = "sjrg7";
  const jobName = "sj9093";
  const functionName = "function588";
  const credential = new DefaultAzureCredential();
  const client = new StreamAnalyticsManagementClient(credential, subscriptionId);
  const result = await client.functions.get(resourceGroupName, jobName, functionName);
  console.log(result);
}

getAnAzureMlFunction().catch(console.error);
