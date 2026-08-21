const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to **Lists the metric values for a resource**. This API used the [default ARM throttling limits](https://learn.microsoft.com/en-us/azure/azure-resource-manager/management/request-limits-and-throttling).
 *
 * @summary **Lists the metric values for a resource**. This API used the [default ARM throttling limits](https://learn.microsoft.com/en-us/azure/azure-resource-manager/management/request-limits-and-throttling).
 * x-ms-original-file: 2024-02-01/GetMetricError.json
 */
async function getMetricWithError() {
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential);
  const result = await client.metrics.list(
    "subscriptions/ac41e21f-afd6-4a79-8070-f01eba278f97/resourceGroups/todking/providers/Microsoft.DocumentDb/databaseAccounts/tk-cosmos-mongo",
    {
      timespan: "2021-06-07T21:51:00Z/2021-06-08T01:51:00Z",
      interval: "FULL",
      metricnames: "MongoRequestsCount,MongoRequests",
      aggregation: "average",
      metricnamespace: "microsoft.documentdb/databaseaccounts",
      autoAdjustTimegrain: true,
      validateDimensions: false,
    },
  );
  console.log(result);
}
