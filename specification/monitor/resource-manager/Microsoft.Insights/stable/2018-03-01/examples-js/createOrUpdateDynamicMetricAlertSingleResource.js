const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update an metric alert definition.
 *
 * @summary Create or update an metric alert definition.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2018-03-01/examples/createOrUpdateDynamicMetricAlertSingleResource.json
 */
async function createOrUpdateADynamicAlertRuleForSingleResource() {
  const subscriptionId =
    process.env["MONITOR_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["MONITOR_RESOURCE_GROUP"] || "gigtest";
  const ruleName = "chiricutin";
  const parameters = {
    description: "This is the description of the rule1",
    actions: [
      {
        actionGroupId:
          "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/gigtest/providers/microsoft.insights/actiongroups/group2",
        webHookProperties: { key11: "value11", key12: "value12" },
      },
    ],
    autoMitigate: true,
    criteria: {
      allOf: [
        {
          name: "High_CPU_80",
          alertSensitivity: "Medium",
          criterionType: "DynamicThresholdCriterion",
          dimensions: [],
          failingPeriods: {
            minFailingPeriodsToAlert: 4,
            numberOfEvaluationPeriods: 4,
          },
          ignoreDataBefore: new Date("2019-04-04T21:00:00.000Z"),
          metricName: "Percentage CPU",
          metricNamespace: "microsoft.compute/virtualmachines",
          operator: "GreaterOrLessThan",
          timeAggregation: "Average",
        },
      ],
      odataType: "Microsoft.Azure.Monitor.MultipleResourceMultipleMetricCriteria",
    },
    enabled: true,
    evaluationFrequency: "PT1M",
    location: "global",
    scopes: [
      "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/gigtest/providers/Microsoft.Compute/virtualMachines/gigwadme",
    ],
    severity: 3,
    tags: {},
    windowSize: "PT15M",
  };
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.metricAlerts.createOrUpdate(resourceGroupName, ruleName, parameters);
  console.log(result);
}
