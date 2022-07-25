const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a classic metric alert rule.
 *
 * @summary Creates or updates a classic metric alert rule.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2016-03-01/examples/createOrUpdateAlertRule.json
 */
async function createOrUpdateAnAlertRule() {
  const subscriptionId = "b67f7fec-69fc-4974-9099-a26bd6ffeda3";
  const resourceGroupName = "Rac46PostSwapRG";
  const ruleName = "chiricutin";
  const parameters = {
    namePropertiesName: "chiricutin",
    description: "Pura Vida",
    actions: [],
    condition: {
      dataSource: {
        metricName: "Requests",
        odataType: "Microsoft.Azure.Management.Insights.Models.RuleMetricDataSource",
        resourceUri:
          "/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/leoalerttest",
      },
      odataType: "Microsoft.Azure.Management.Insights.Models.ThresholdRuleCondition",
      operator: "GreaterThan",
      threshold: 3,
      timeAggregation: "Total",
      windowSize: "PT5M",
    },
    isEnabled: true,
    location: "West US",
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.alertRules.createOrUpdate(resourceGroupName, ruleName, parameters);
  console.log(result);
}

createOrUpdateAnAlertRule().catch(console.error);
