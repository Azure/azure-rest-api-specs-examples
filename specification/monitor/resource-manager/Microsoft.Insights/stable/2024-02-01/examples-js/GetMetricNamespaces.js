const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the metric namespaces for the resource.
 *
 * @summary Lists the metric namespaces for the resource.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2024-02-01/examples/GetMetricNamespaces.json
 */
async function getMetricNamespacesWithoutFilter() {
  const resourceUri =
    "subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill";
  const startTime = "2020-08-31T15:53:00Z";
  const options = { startTime };
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential);
  const resArray = new Array();
  for await (let item of client.metricNamespaces.list(resourceUri, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}
