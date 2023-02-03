const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of incidents associated to an alert rule
 *
 * @summary Gets a list of incidents associated to an alert rule
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2016-03-01/examples/listAlertRuleIncidents.json
 */
async function listAlertRuleIncidents() {
  const subscriptionId =
    process.env["MONITOR_SUBSCRIPTION_ID"] || "b67f7fec-69fc-4974-9099-a26bd6ffeda3";
  const resourceGroupName = process.env["MONITOR_RESOURCE_GROUP"] || "Rac46PostSwapRG";
  const ruleName = "myRuleName";
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.alertRuleIncidents.listByAlertRule(resourceGroupName, ruleName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
