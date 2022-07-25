const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update an metric alert definition.
 *
 * @summary Create or update an metric alert definition.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2018-03-01/examples/createOrUpdateMetricAlertWithDimensions.json
 */
async function createOrUpdateAnAlertRulesWithDimensions() {
  const subscriptionId = "14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7";
  const resourceGroupName = "gigtest";
  const ruleName = "MetricAlertOnMultipleDimensions";
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
    location: "global",
    scopes: [
      "/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourceGroups/gigtest/providers/Microsoft.KeyVault/vaults/keyVaultResource",
    ],
    severity: 3,
    tags: {},
    windowSize: "P1D",
  };
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.metricAlerts.createOrUpdate(resourceGroupName, ruleName, parameters);
  console.log(result);
}

createOrUpdateAnAlertRulesWithDimensions().catch(console.error);
