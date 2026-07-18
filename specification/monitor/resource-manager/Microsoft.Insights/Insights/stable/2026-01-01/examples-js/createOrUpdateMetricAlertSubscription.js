const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update an metric alert definition.
 *
 * @summary create or update an metric alert definition.
 * x-ms-original-file: 2026-01-01/createOrUpdateMetricAlertSubscription.json
 */
async function createOrUpdateAnAlertRuleOnSubscription() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7";
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.metricAlerts.createOrUpdate(
    "gigtest",
    "MetricAlertAtSubscriptionLevel",
    {
      location: "global",
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
            metricName: "Percentage CPU",
            metricNamespace: "microsoft.compute/virtualmachines",
            operator: "GreaterThan",
            threshold: 80.5,
            timeAggregation: "Average",
          },
        ],
        odataType: "Microsoft.Azure.Monitor.MultipleResourceMultipleMetricCriteria",
      },
      enabled: true,
      evaluationFrequency: "PT1M",
      scopes: ["/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7"],
      severity: 3,
      targetResourceRegion: "southcentralus",
      targetResourceType: "Microsoft.Compute/virtualMachines",
      windowSize: "PT15M",
      tags: {},
    },
  );
  console.log(result);
}
