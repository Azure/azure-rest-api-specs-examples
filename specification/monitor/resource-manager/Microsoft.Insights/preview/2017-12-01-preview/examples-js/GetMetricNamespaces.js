const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the metric namespaces for the resource.
 *
 * @summary Lists the metric namespaces for the resource.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/preview/2017-12-01-preview/examples/GetMetricNamespaces.json
 */
async function getMetricNamespacesWithoutFilter() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceUri =
    "subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill";
  const startTime = "2020-08-31T15:53:00Z";
  const options = { startTime };
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.metricNamespaces.list(resourceUri, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getMetricNamespacesWithoutFilter().catch(console.error);
