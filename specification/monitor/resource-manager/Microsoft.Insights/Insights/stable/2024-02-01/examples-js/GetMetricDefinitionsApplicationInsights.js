const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to lists the metric definitions for the resource.
 *
 * @summary lists the metric definitions for the resource.
 * x-ms-original-file: 2024-02-01/GetMetricDefinitionsApplicationInsights.json
 */
async function getApplicationInsightsMetricDefinitionsWithoutFilter() {
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential);
  const resArray = new Array();
  for await (const item of client.metricDefinitions.list(
    "subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill/providers/microsoft.insights/metricdefinitions",
    { metricnamespace: "microsoft.insights/components" },
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
