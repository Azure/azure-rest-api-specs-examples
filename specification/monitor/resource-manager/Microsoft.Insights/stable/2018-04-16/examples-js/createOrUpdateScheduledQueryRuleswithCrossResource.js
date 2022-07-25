const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an log search rule.
 *
 * @summary Creates or updates an log search rule.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2018-04-16/examples/createOrUpdateScheduledQueryRuleswithCrossResource.json
 */
async function createOrUpdateRuleAlertingActionWithCrossResource() {
  const subscriptionId = "b67f7fec-69fc-4974-9099-a26bd6ffeda3";
  const resourceGroupName = "Rac46PostSwapRG";
  const ruleName = "SampleCrossResourceAlert";
  const parameters = {
    description: "Sample Cross Resource alert",
    action: {
      aznsAction: {
        actionGroup: [
          "/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/Rac46PostSwapRG/providers/microsoft.insights/actiongroups/test-ag",
        ],
        emailSubject: "Cross Resource Mail!!",
      },
      odataType:
        "Microsoft.WindowsAzure.Management.Monitoring.Alerts.Models.Microsoft.AppInsights.Nexus.DataContracts.Resources.ScheduledQueryRules.AlertingAction",
      severity: "3",
      trigger: { threshold: 5000, thresholdOperator: "GreaterThan" },
    },
    enabled: "true",
    location: "eastus",
    schedule: { frequencyInMinutes: 60, timeWindowInMinutes: 60 },
    source: {
      authorizedResources: [
        "/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/Rac46PostSwapRG/providers/Microsoft.OperationalInsights/workspaces/sampleWorkspace",
        "/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/Rac46PostSwapRG/providers/microsoft.insights/components/sampleAI",
      ],
      dataSourceId:
        "/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/Rac46PostSwapRG/providers/microsoft.insights/components/sampleAI",
      query: 'union requests, workspace("sampleWorkspace").Update',
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

createOrUpdateRuleAlertingActionWithCrossResource().catch(console.error);
