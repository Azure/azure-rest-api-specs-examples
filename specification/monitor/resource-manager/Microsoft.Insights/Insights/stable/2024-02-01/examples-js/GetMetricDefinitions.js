const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to lists the metric definitions for the resource.
 *
 * @summary lists the metric definitions for the resource.
 * x-ms-original-file: 2024-02-01/GetMetricDefinitions.json
 */
async function getMetricDefinitionsWithoutFilter() {
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential);
  const resArray = new Array();
  for await (const item of client.metricDefinitions.list(
    "subscriptions/07c0b09d-9f69-4e6e-8d05-f59f67299cb2/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/alertruleTest/providers/microsoft.insights/metricDefinitions",
    { metricnamespace: "Microsoft.Web/sites" },
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
