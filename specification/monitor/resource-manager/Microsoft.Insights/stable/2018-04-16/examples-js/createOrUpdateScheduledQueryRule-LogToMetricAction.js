const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an log search rule.
 *
 * @summary Creates or updates an log search rule.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2018-04-16/examples/createOrUpdateScheduledQueryRule-LogToMetricAction.json
 */
async function createOrUpdateRuleLogToMetricAction() {
  const subscriptionId =
    process.env["MONITOR_SUBSCRIPTION_ID"] || "af52d502-a447-4bc6-8cb7-4780fbb00490";
  const resourceGroupName = process.env["MONITOR_RESOURCE_GROUP"] || "alertsweu";
  const ruleName = "logtometricfoo";
  const parameters = {
    description: "log to metric description",
    action: {
      criteria: [{ dimensions: [], metricName: "Average_% Idle Time" }],
      odataType:
        "Microsoft.WindowsAzure.Management.Monitoring.Alerts.Models.Microsoft.AppInsights.Nexus.DataContracts.Resources.ScheduledQueryRules.LogToMetricAction",
    },
    enabled: "true",
    location: "West Europe",
    source: {
      dataSourceId:
        "/subscriptions/af52d502-a447-4bc6-8cb7-4780fbb00490/resourceGroups/alertsweu/providers/Microsoft.OperationalInsights/workspaces/alertsweu",
    },
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.scheduledQueryRules.createOrUpdate(
    resourceGroupName,
    ruleName,
    parameters
  );
  console.log(result);
}
