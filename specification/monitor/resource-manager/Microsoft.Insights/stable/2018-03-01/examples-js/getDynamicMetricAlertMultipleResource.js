const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve an alert rule definition.
 *
 * @summary Retrieve an alert rule definition.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2018-03-01/examples/getDynamicMetricAlertMultipleResource.json
 */
async function getADynamicAlertRuleForMultipleResources() {
  const subscriptionId =
    process.env["MONITOR_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["MONITOR_RESOURCE_GROUP"] || "gigtest";
  const ruleName = "MetricAlertOnMultipleResources";
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.metricAlerts.get(resourceGroupName, ruleName);
  console.log(result);
}
