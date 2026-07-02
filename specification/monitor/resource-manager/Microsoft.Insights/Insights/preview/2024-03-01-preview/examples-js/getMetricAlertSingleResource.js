const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to retrieve an alert rule definition.
 *
 * @summary retrieve an alert rule definition.
 * x-ms-original-file: 2024-03-01-preview/getMetricAlertSingleResource.json
 */
async function getAnAlertRuleForSingleResource() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7";
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.metricAlerts.get("gigtest", "chiricutin");
  console.log(result);
}
