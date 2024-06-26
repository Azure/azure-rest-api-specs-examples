const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Test the Stream Analytics output.
 *
 * @summary Test the Stream Analytics output.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Subscription_TestOutput.json
 */
async function testTheStreamAnalyticsOutput() {
  const subscriptionId =
    process.env["STREAMANALYTICS_SUBSCRIPTION_ID"] || "56b5e0a9-b645-407d-99b0-c64f86013e3d";
  const location = "West US";
  const testOutput = {
    output: {
      datasource: {
        type: "Microsoft.Storage/Blob",
        container: "state",
        dateFormat: "yyyy/MM/dd",
        pathPattern: "{date}/{time}",
        storageAccounts: [{ accountKey: "accountKey==", accountName: "someAccountName" }],
        timeFormat: "HH",
      },
      serialization: { type: "Csv", encoding: "UTF8", fieldDelimiter: "," },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new StreamAnalyticsManagementClient(credential, subscriptionId);
  const result = await client.subscriptions.beginTestOutputAndWait(location, testOutput);
  console.log(result);
}
