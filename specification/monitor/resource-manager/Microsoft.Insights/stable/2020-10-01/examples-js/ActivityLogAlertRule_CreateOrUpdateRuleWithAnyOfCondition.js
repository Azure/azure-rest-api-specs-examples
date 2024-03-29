const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a new Activity Log Alert rule or update an existing one.
 *
 * @summary Create a new Activity Log Alert rule or update an existing one.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2020-10-01/examples/ActivityLogAlertRule_CreateOrUpdateRuleWithAnyOfCondition.json
 */
async function createOrUpdateAnActivityLogAlertRuleWithAnyOfCondition() {
  const subscriptionId =
    process.env["MONITOR_SUBSCRIPTION_ID"] || "187f412d-1758-44d9-b052-169e2564721d";
  const resourceGroupName = process.env["MONITOR_RESOURCE_GROUP"] || "MyResourceGroup";
  const activityLogAlertName = "SampleActivityLogAlertRuleWithAnyOfCondition";
  const activityLogAlertRule = {
    description: "Description of sample Activity Log Alert rule with 'anyOf' condition.",
    actions: {
      actionGroups: [
        {
          actionGroupId:
            "/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/MyResourceGroup/providers/Microsoft.Insights/actionGroups/SampleActionGroup",
          webhookProperties: { sampleWebhookProperty: "SamplePropertyValue" },
        },
      ],
    },
    condition: {
      allOf: [
        { equals: "ServiceHealth", field: "category" },
        {
          anyOf: [
            { equals: "Incident", field: "properties.incidentType" },
            { equals: "Maintenance", field: "properties.incidentType" },
          ],
        },
      ],
    },
    enabled: true,
    location: "Global",
    scopes: ["subscriptions/187f412d-1758-44d9-b052-169e2564721d"],
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.activityLogAlerts.createOrUpdate(
    resourceGroupName,
    activityLogAlertName,
    activityLogAlertRule,
  );
  console.log(result);
}
