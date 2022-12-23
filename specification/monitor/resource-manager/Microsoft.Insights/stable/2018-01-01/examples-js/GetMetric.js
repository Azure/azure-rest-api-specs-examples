const { MonitorClient } = require("@azure/arm-monitor-profile-2020-09-01-hybrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to **Lists the metric values for a resource**.
 *
 * @summary **Lists the metric values for a resource**.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2018-01-01/examples/GetMetric.json
 */
async function getMetricForData() {
  const resourceUri =
    "subscriptions/b324c52b-4073-4807-93af-e07d289c093e/resourceGroups/test/providers/Microsoft.Storage/storageAccounts/larryshoebox/blobServices/default";
  const timespan = "2017-04-14T02:20:00Z/2017-04-14T04:20:00Z";
  const interval = "PT1M";
  const aggregation = "Average,count";
  const top = 3;
  const orderby = "Average asc";
  const filter = "BlobType eq '*'";
  const metricnamespace = "Microsoft.Storage/storageAccounts/blobServices";
  const options = {
    timespan,
    interval,
    aggregation,
    top,
    orderby,
    filter,
    metricnamespace,
  };
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential);
  const result = await client.metrics.list(resourceUri, options);
  console.log(result);
}
