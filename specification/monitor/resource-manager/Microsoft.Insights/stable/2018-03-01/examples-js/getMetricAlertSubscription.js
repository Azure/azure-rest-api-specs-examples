const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve an alert rule definition.
 *
 * @summary Retrieve an alert rule definition.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2018-03-01/examples/getMetricAlertSubscription.json
 */
async function getAnAlertRuleOnSubscription() {
  const subscriptionId = "14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7";
  const resourceGroupName = "gigtest";
  const ruleName = "MetricAlertAtSubscriptionLevel";
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.metricAlerts.get(resourceGroupName, ruleName);
  console.log(result);
}

getAnAlertRuleOnSubscription().catch(console.error);
