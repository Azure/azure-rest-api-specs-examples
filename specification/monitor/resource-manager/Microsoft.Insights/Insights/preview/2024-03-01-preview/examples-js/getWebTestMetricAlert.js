const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to retrieve an alert rule definition.
 *
 * @summary retrieve an alert rule definition.
 * x-ms-original-file: 2024-03-01-preview/getWebTestMetricAlert.json
 */
async function getAWebTestAlertRule() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "12345678-1234-1234-1234-123456789101";
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.metricAlerts.get("rg-example", "webtest-name-example");
  console.log(result);
}
