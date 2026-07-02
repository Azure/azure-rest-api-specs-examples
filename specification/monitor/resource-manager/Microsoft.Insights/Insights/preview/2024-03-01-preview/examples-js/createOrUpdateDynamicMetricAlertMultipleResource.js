const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update an metric alert definition.
 *
 * @summary create or update an metric alert definition.
 * x-ms-original-file: 2024-03-01-preview/createOrUpdateDynamicMetricAlertMultipleResource.json
 */
async function createOrUpdateADynamicAlertRuleForMultipleResources() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.metricAlerts.createOrUpdate(
    "gigtest",
    "MetricAlertOnMultipleResources",
    {
      location: "global",
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
            failingPeriods: { minFailingPeriodsToAlert: 4, numberOfEvaluationPeriods: 4 },
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
      scopes: [
        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/gigtest/providers/Microsoft.Compute/virtualMachines/gigwadme1",
        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/gigtest/providers/Microsoft.Compute/virtualMachines/gigwadme2",
      ],
      severity: 3,
      targetResourceRegion: "southcentralus",
      targetResourceType: "Microsoft.Compute/virtualMachines",
      windowSize: "PT15M",
      tags: {},
    },
  );
  console.log(result);
}
