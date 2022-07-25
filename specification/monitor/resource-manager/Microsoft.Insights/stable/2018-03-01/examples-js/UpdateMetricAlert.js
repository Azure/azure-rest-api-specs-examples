const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update an metric alert definition.
 *
 * @summary Update an metric alert definition.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2018-03-01/examples/UpdateMetricAlert.json
 */
async function createOrUpdateAnAlertRule() {
  const subscriptionId = "14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7";
  const resourceGroupName = "gigtest";
  const ruleName = "chiricutin";
  const parameters = {
    description: "This is the description of the rule1",
    actions: [
      {
        actionGroupId:
          "/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourcegroups/gigtest/providers/microsoft.insights/actiongroups/group2",
        webHookProperties: { key11: "value11", key12: "value12" },
      },
    ],
    autoMitigate: true,
    criteria: {
      allOf: [
        {
          name: "High_CPU_80",
          criterionType: "StaticThresholdCriterion",
          dimensions: [],
          metricName: "Processor(_Total)% Processor Time",
          operator: "GreaterThan",
          threshold: 80.5,
          timeAggregation: "Average",
        },
      ],
      odataType: "Microsoft.Azure.Monitor.SingleResourceMultipleMetricCriteria",
    },
    enabled: true,
    evaluationFrequency: "Pt1m",
    scopes: [
      "/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourceGroups/gigtest/providers/Microsoft.Compute/virtualMachines/gigwadme",
    ],
    severity: 3,
    tags: {},
    windowSize: "Pt15m",
  };
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.metricAlerts.update(resourceGroupName, ruleName, parameters);
  console.log(result);
}

createOrUpdateAnAlertRule().catch(console.error);
