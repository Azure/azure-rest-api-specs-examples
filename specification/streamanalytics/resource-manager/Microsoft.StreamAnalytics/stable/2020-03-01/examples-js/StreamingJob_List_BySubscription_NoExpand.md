```javascript
const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the streaming jobs in the given subscription.
 *
 * @summary Lists all of the streaming jobs in the given subscription.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/StreamingJob_List_BySubscription_NoExpand.json
 */
async function listAllStreamingJobsInASubscriptionAndDoNotUseTheExpandODataQueryParameter() {
  const subscriptionId = "56b5e0a9-b645-407d-99b0-c64f86013e3d";
  const credential = new DefaultAzureCredential();
  const client = new StreamAnalyticsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.streamingJobs.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAllStreamingJobsInASubscriptionAndDoNotUseTheExpandODataQueryParameter().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-streamanalytics_4.0.1/sdk/streamanalytics/arm-streamanalytics/README.md) on how to add the SDK to your project and authenticate.
