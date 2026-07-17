const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update an metric alert definition.
 *
 * @summary create or update an metric alert definition.
 * x-ms-original-file: 2026-01-01/createOrUpdateMetricAlertWithDimensions.json
 */
async function createOrUpdateAnAlertRulesWithDimensions() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7";
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.metricAlerts.createOrUpdate(
    "gigtest",
    "MetricAlertOnMultipleDimensions",
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
            name: "Metric1",
            criterionType: "StaticThresholdCriterion",
            dimensions: [
              { name: "ActivityName", operator: "Include", values: ["*"] },
              { name: "StatusCode", operator: "Include", values: ["200"] },
            ],
            metricName: "Availability",
            metricNamespace: "Microsoft.KeyVault/vaults",
            operator: "GreaterThan",
            threshold: 55,
            timeAggregation: "Average",
          },
        ],
        odataType: "Microsoft.Azure.Monitor.MultipleResourceMultipleMetricCriteria",
      },
      enabled: true,
      evaluationFrequency: "PT1H",
      scopes: [
        "/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourceGroups/gigtest/providers/Microsoft.KeyVault/vaults/keyVaultResource",
      ],
      severity: 3,
      windowSize: "P1D",
      tags: {},
    },
  );
  console.log(result);
}
