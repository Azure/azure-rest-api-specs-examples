const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve alert rule definitions in a subscription.
 *
 * @summary Retrieve alert rule definitions in a subscription.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2018-03-01/examples/listMetricAlert.json
 */
async function listMetricAlertRules() {
  const subscriptionId =
    process.env["MONITOR_SUBSCRIPTION_ID"] || "14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7";
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.metricAlerts.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
