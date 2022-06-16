const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Tests if the information provided for a function is valid. This can range from testing the connection to the underlying web service behind the function or making sure the function code provided is syntactically correct.
 *
 * @summary Tests if the information provided for a function is valid. This can range from testing the connection to the underlying web service behind the function or making sure the function code provided is syntactically correct.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Function_Test_JavaScript.json
 */
async function testTheConnectionForAJavaScriptFunction() {
  const subscriptionId = "56b5e0a9-b645-407d-99b0-c64f86013e3d";
  const resourceGroupName = "sjrg1637";
  const jobName = "sj8653";
  const functionName = "function8197";
  const credential = new DefaultAzureCredential();
  const client = new StreamAnalyticsManagementClient(credential, subscriptionId);
  const result = await client.functions.beginTestAndWait(resourceGroupName, jobName, functionName);
  console.log(result);
}

testTheConnectionForAJavaScriptFunction().catch(console.error);
