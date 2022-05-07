Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-streamanalytics_4.0.1/sdk/streamanalytics/arm-streamanalytics/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Stops a running streaming job. This will cause a running streaming job to stop processing input events and producing output.
 *
 * @summary Stops a running streaming job. This will cause a running streaming job to stop processing input events and producing output.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/StreamingJob_Stop.json
 */
async function stopAStreamingJob() {
  const subscriptionId = "56b5e0a9-b645-407d-99b0-c64f86013e3d";
  const resourceGroupName = "sjrg6936";
  const jobName = "sj59";
  const credential = new DefaultAzureCredential();
  const client = new StreamAnalyticsManagementClient(credential, subscriptionId);
  const result = await client.streamingJobs.beginStopAndWait(resourceGroupName, jobName);
  console.log(result);
}

stopAStreamingJob().catch(console.error);
```
