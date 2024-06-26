const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets details about the specified function.
 *
 * @summary Gets details about the specified function.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Function_Get_AzureML.json
 */
async function getAnAzureMlFunction() {
  const subscriptionId =
    process.env["STREAMANALYTICS_SUBSCRIPTION_ID"] || "56b5e0a9-b645-407d-99b0-c64f86013e3d";
  const resourceGroupName = process.env["STREAMANALYTICS_RESOURCE_GROUP"] || "sjrg7";
  const jobName = "sj9093";
  const functionName = "function588";
  const credential = new DefaultAzureCredential();
  const client = new StreamAnalyticsManagementClient(credential, subscriptionId);
  const result = await client.functions.get(resourceGroupName, jobName, functionName);
  console.log(result);
}
