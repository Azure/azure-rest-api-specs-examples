const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the metric definitions for the resource.
 *
 * @summary Lists the metric definitions for the resource.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2024-02-01/examples/GetMetricDefinitionsMetricClass.json
 */
async function getStorageCacheMetricDefinitionsWithMetricClass() {
  const resourceUri =
    "subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache";
  const metricnamespace = "microsoft.storagecache/caches";
  const options = { metricnamespace };
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential);
  const resArray = new Array();
  for await (let item of client.metricDefinitions.list(resourceUri, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}
