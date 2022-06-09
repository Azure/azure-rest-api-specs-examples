Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-compute_19.0.0/sdk/compute/arm-compute/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Export logs that show total throttled Api requests for this subscription in the given time window.
 *
 * @summary Export logs that show total throttled Api requests for this subscription in the given time window.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/logAnalyticExamples/LogAnalytics_ThrottledRequests.json
 */
async function exportLogsWhichContainAllThrottledApiRequestsMadeToComputeResourceProviderWithinTheGivenTimePeriod() {
  const subscriptionId = "{subscription-id}";
  const location = "westus";
  const parameters = {
    blobContainerSasUri: "https://somesasuri",
    fromTime: new Date("2018-01-21T01:54:06.862601Z"),
    groupByClientApplicationId: false,
    groupByOperationName: true,
    groupByResourceName: false,
    groupByUserAgent: false,
    toTime: new Date("2018-01-23T01:54:06.862601Z"),
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.logAnalytics.beginExportThrottledRequestsAndWait(
    location,
    parameters
  );
  console.log(result);
}

exportLogsWhichContainAllThrottledApiRequestsMadeToComputeResourceProviderWithinTheGivenTimePeriod().catch(
  console.error
);
```
