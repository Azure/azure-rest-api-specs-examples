const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve an alert rule definition.
 *
 * @summary Retrieve an alert rule definition.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2018-03-01/examples/getWebTestMetricAlert.json
 */
async function getAWebTestAlertRule() {
  const subscriptionId =
    process.env["MONITOR_SUBSCRIPTION_ID"] || "12345678-1234-1234-1234-123456789101";
  const resourceGroupName = process.env["MONITOR_RESOURCE_GROUP"] || "rg-example";
  const ruleName = "webtest-name-example";
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.metricAlerts.get(resourceGroupName, ruleName);
  console.log(result);
}
