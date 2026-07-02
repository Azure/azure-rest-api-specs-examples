const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to **Lists the metric values for a resource**. This API used the [default ARM throttling limits](https://learn.microsoft.com/en-us/azure/azure-resource-manager/management/request-limits-and-throttling).
 *
 * @summary **Lists the metric values for a resource**. This API used the [default ARM throttling limits](https://learn.microsoft.com/en-us/azure/azure-resource-manager/management/request-limits-and-throttling).
 * x-ms-original-file: 2024-02-01/GetMetricMetadata.json
 */
async function getMetricForMetadata() {
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential);
  const result = await client.metrics.list(
    "subscriptions/1f3fa6d2-851c-4a91-9087-1a050f3a9c38/resourceGroups/todking/providers/Microsoft.Storage/storageAccounts/tkfileserv/blobServices/default",
    {
      timespan: "2017-04-14T02:20:00Z/2017-04-14T04:20:00Z",
      filter: "Tier eq '*'",
      metricnamespace: "Microsoft.Storage/storageAccounts/blobServices",
    },
  );
  console.log(result);
}
