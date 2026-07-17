const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to **Lists the metric values for a resource**. This API used the [default ARM throttling limits](https://learn.microsoft.com/azure/azure-resource-manager/management/request-limits-and-throttling).
 *
 * @summary **Lists the metric values for a resource**. This API used the [default ARM throttling limits](https://learn.microsoft.com/azure/azure-resource-manager/management/request-limits-and-throttling).
 * x-ms-original-file: 2024-02-01/GetMetric.json
 */
async function getMetricForData() {
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential);
  const result = await client.metrics.list(
    "subscriptions/1f3fa6d2-851c-4a91-9087-1a050f3a9c38/resourceGroups/todking/providers/Microsoft.Storage/storageAccounts/tkfileserv/blobServices/default",
    {
      timespan: "2021-04-20T09:00:00.000Z/2021-04-20T14:00:00.000Z",
      interval: "PT6H",
      metricnames: "BlobCount,BlobCapacity",
      aggregation: "average,minimum,maximum",
      top: 5,
      orderby: "average asc",
      filter: "Tier eq '*'",
      metricnamespace: "Microsoft.Storage/storageAccounts/blobServices",
      autoAdjustTimegrain: true,
      validateDimensions: false,
    },
  );
  console.log(result);
}
