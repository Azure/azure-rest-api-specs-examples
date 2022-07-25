const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an log search rule.
 *
 * @summary Creates or updates an log search rule.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2018-04-16/examples/createOrUpdateScheduledQueryRules.json
 */
async function createOrUpdateRuleAlertingAction() {
  const subscriptionId = "b67f7fec-69fc-4974-9099-a26bd6ffeda3";
  const resourceGroupName = "Rac46PostSwapRG";
  const ruleName = "logalertfoo";
  const parameters = {
    description: "log alert description",
    action: {
      aznsAction: {
        actionGroup: [],
        customWebhookPayload: "{}",
        emailSubject: "Email Header",
      },
      odataType:
        "Microsoft.WindowsAzure.Management.Monitoring.Alerts.Models.Microsoft.AppInsights.Nexus.DataContracts.Resources.ScheduledQueryRules.AlertingAction",
      severity: "1",
      trigger: {
        metricTrigger: {
          metricColumn: "Computer",
          metricTriggerType: "Consecutive",
          threshold: 5,
          thresholdOperator: "GreaterThan",
        },
        threshold: 3,
        thresholdOperator: "GreaterThan",
      },
    },
    enabled: "true",
    location: "eastus",
    schedule: { frequencyInMinutes: 15, timeWindowInMinutes: 15 },
    source: {
      dataSourceId:
        "/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/Rac46PostSwapRG/providers/Microsoft.OperationalInsights/workspaces/sampleWorkspace",
      query: "Heartbeat | summarize AggregatedValue = count() by bin(TimeGenerated, 5m)",
      queryType: "ResultCount",
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

createOrUpdateRuleAlertingAction().catch(console.error);
