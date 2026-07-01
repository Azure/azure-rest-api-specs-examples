const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to lists the metric definitions for the resource.
 *
 * @summary lists the metric definitions for the resource.
 * x-ms-original-file: 2024-02-01/GetMetricDefinitionsMetricClass.json
 */
async function getStorageCacheMetricDefinitionsWithMetricClass() {
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential);
  const resArray = new Array();
  for await (const item of client.metricDefinitions.list(
    "subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache",
    { metricnamespace: "microsoft.storagecache/caches" },
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
