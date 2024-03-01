const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Tests whether an output’s datasource is reachable and usable by the Azure Stream Analytics service.
 *
 * @summary Tests whether an output’s datasource is reachable and usable by the Azure Stream Analytics service.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Output_Test.json
 */
async function testTheConnectionForAnOutput() {
  const subscriptionId =
    process.env["STREAMANALYTICS_SUBSCRIPTION_ID"] || "56b5e0a9-b645-407d-99b0-c64f86013e3d";
  const resourceGroupName = process.env["STREAMANALYTICS_RESOURCE_GROUP"] || "sjrg2157";
  const jobName = "sj6458";
  const outputName = "output1755";
  const credential = new DefaultAzureCredential();
  const client = new StreamAnalyticsManagementClient(credential, subscriptionId);
  const result = await client.outputs.beginTestAndWait(resourceGroupName, jobName, outputName);
  console.log(result);
}
