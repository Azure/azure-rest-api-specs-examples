const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a classic metric alert rule
 *
 * @summary Deletes a classic metric alert rule
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2016-03-01/examples/deleteAlertRule.json
 */
async function deleteAnAlertRule() {
  const subscriptionId = "b67f7fec-69fc-4974-9099-a26bd6ffeda3";
  const resourceGroupName = "Rac46PostSwapRG";
  const ruleName = "chiricutin";
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.alertRules.delete(resourceGroupName, ruleName);
  console.log(result);
}

deleteAnAlertRule().catch(console.error);
