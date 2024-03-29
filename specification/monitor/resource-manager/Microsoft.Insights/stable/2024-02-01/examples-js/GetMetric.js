const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to **Lists the metric values for a resource**.
 *
 * @summary **Lists the metric values for a resource**.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2024-02-01/examples/GetMetric.json
 */
async function getMetricForData() {
  const resourceUri =
    "subscriptions/1f3fa6d2-851c-4a91-9087-1a050f3a9c38/resourceGroups/todking/providers/Microsoft.Storage/storageAccounts/tkfileserv/blobServices/default";
  const timespan = "2021-04-20T09:00:00.000Z/2021-04-20T14:00:00.000Z";
  const interval = "PT6H";
  const metricnames = "BlobCount,BlobCapacity";
  const aggregation = "average,minimum,maximum";
  const top = 5;
  const orderby = "average asc";
  const filter = "Tier eq '*'";
  const metricnamespace = "Microsoft.Storage/storageAccounts/blobServices";
  const autoAdjustTimegrain = true;
  const validateDimensions = false;
  const options = {
    timespan,
    interval,
    metricnames,
    aggregation,
    top,
    orderby,
    filter,
    metricnamespace,
    autoAdjustTimegrain,
    validateDimensions,
  };
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential);
  const result = await client.metricsOperations.list(resourceUri, options);
  console.log(result);
}
