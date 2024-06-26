const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets an incident associated to an alert rule
 *
 * @summary Gets an incident associated to an alert rule
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2016-03-01/examples/getAlertRuleIncident.json
 */
async function getASingleAlertRuleIncident() {
  const subscriptionId =
    process.env["MONITOR_SUBSCRIPTION_ID"] || "b67f7fec-69fc-4974-9099-a26bd6ffeda3";
  const resourceGroupName = process.env["MONITOR_RESOURCE_GROUP"] || "Rac46PostSwapRG";
  const ruleName = "myRuleName";
  const incidentName = "Website_started";
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.alertRuleIncidents.get(resourceGroupName, ruleName, incidentName);
  console.log(result);
}
