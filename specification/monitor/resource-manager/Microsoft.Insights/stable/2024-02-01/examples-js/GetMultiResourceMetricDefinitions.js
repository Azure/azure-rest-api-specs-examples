const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the metric definitions for the subscription.
 *
 * @summary Lists the metric definitions for the subscription.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2024-02-01/examples/GetMultiResourceMetricDefinitions.json
 */
async function getSubscriptionLevelMetricDefinitionsWithoutFilter() {
  const subscriptionId =
    process.env["MONITOR_SUBSCRIPTION_ID"] || "92d2a2d8-b514-432d-8cc9-a5f9272630d5";
  const region = "westus2";
  const metricnamespace = "microsoft.compute/virtualmachines";
  const options = {
    metricnamespace,
  };
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.metricDefinitions.listAtSubscriptionScope(region, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}
