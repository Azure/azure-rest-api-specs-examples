const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update an metric alert definition.
 *
 * @summary Create or update an metric alert definition.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2018-03-01/examples/createOrUpdateMetricAlertResourceGroup.json
 */
async function createOrUpdateAnAlertRuleOnResourceGroupS() {
  const subscriptionId = "14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7";
  const resourceGroupName = "gigtest1";
  const ruleName = "MetricAlertAtResourceGroupLevel";
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
    location: "global",
    scopes: [
      "/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourceGroups/gigtest1",
      "/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourceGroups/gigtest2",
    ],
    severity: 3,
    tags: {},
    targetResourceRegion: "southcentralus",
    targetResourceType: "Microsoft.Compute/virtualMachines",
    windowSize: "PT15M",
  };
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.metricAlerts.createOrUpdate(resourceGroupName, ruleName, parameters);
  console.log(result);
}

createOrUpdateAnAlertRuleOnResourceGroupS().catch(console.error);
