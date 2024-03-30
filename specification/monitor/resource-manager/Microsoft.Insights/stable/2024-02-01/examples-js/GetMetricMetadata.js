const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to **Lists the metric values for a resource**.
 *
 * @summary **Lists the metric values for a resource**.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2024-02-01/examples/GetMetricMetadata.json
 */
async function getMetricForMetadata() {
  const resourceUri =
    "subscriptions/1f3fa6d2-851c-4a91-9087-1a050f3a9c38/resourceGroups/todking/providers/Microsoft.Storage/storageAccounts/tkfileserv/blobServices/default";
  const timespan = "2017-04-14T02:20:00Z/2017-04-14T04:20:00Z";
  const filter = "Tier eq '*'";
  const metricnamespace = "Microsoft.Storage/storageAccounts/blobServices";
  const options = {
    timespan,
    filter,
    metricnamespace,
  };
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential);
  const result = await client.metricsOperations.list(resourceUri, options);
  console.log(result);
}
