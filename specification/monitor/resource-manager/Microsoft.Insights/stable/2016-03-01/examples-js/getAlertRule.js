const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a classic metric alert rule
 *
 * @summary Gets a classic metric alert rule
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2016-03-01/examples/getAlertRule.json
 */
async function getAnAlertRule() {
  const subscriptionId = "b67f7fec-69fc-4974-9099-a26bd6ffeda3";
  const resourceGroupName = "Rac46PostSwapRG";
  const ruleName = "chiricutin";
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.alertRules.get(resourceGroupName, ruleName);
  console.log(result);
}

getAnAlertRule().catch(console.error);
