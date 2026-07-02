const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update an metric alert definition.
 *
 * @summary create or update an metric alert definition.
 * x-ms-original-file: 2024-03-01-preview/createOrUpdateMetricAlertQueryDT.json
 */
async function createOrUpdateAQueryBasedAlertRuleWithDynamicThreshold() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7";
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.metricAlerts.createOrUpdate("gigtest", "chiricutin", {
    identity: {
      type: "UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/2f1a501a-6e1d-4f37-a445-462d7f8a563d/resourceGroups/AdisTest/providers/Microsoft.ManagedIdentity/userAssignedIdentities/msi-test-euap":
          {},
      },
    },
    location: "eastus",
    description: "This is the description of the rule1",
    actionProperties: { "Email.Sujbect": "my custom email subject" },
    actions: [
      {
        actionGroupId:
          "/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourcegroups/gigtest/providers/microsoft.insights/actiongroups/group2",
      },
    ],
    criteria: {
      allOf: [
        {
          name: "Metric1",
          alertSensitivity: "Medium",
          criterionType: "DynamicThresholdCriterion",
          ignoreDataBefore: new Date("2019-04-04T21:00:00.000Z"),
          operator: "LessThan",
          query: 'avg({"system.cpu.utilization"})',
        },
      ],
      failingPeriods: { for: "PT5M" },
      odataType: "Microsoft.Azure.Monitor.PromQLCriteria",
    },
    customProperties: { key11: "value11", key12: "value12" },
    enabled: true,
    evaluationFrequency: "PT1M",
    resolveConfiguration: { autoResolved: true, timeToResolve: "PT10M" },
    scopes: [
      "/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourceGroups/gigtest/providers/microsoft.monitor/accounts/gigwadme",
    ],
    severity: 3,
    tags: {},
  });
  console.log(result);
}
