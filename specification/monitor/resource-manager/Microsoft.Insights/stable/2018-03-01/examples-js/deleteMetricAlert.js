const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete an alert rule definition.
 *
 * @summary Delete an alert rule definition.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2018-03-01/examples/deleteMetricAlert.json
 */
async function deleteAnAlertRule() {
  const subscriptionId = "14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7";
  const resourceGroupName = "gigtest";
  const ruleName = "chiricutin";
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.metricAlerts.delete(resourceGroupName, ruleName);
  console.log(result);
}

deleteAnAlertRule().catch(console.error);
