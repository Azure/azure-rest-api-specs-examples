const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a new Activity Log Alert rule or update an existing one.
 *
 * @summary Create a new Activity Log Alert rule or update an existing one.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2020-10-01/examples/ActivityLogAlertRule_CreateOrUpdate.json
 */
async function createOrUpdateAnActivityLogAlertRule() {
  const subscriptionId = "187f412d-1758-44d9-b052-169e2564721d";
  const resourceGroupName = "MyResourceGroup";
  const activityLogAlertName = "SampleActivityLogAlertRule";
  const activityLogAlertRule = {
    description: "Description of sample Activity Log Alert rule.",
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
        { equals: "Administrative", field: "category" },
        { equals: "Error", field: "level" },
      ],
    },
    enabled: true,
    location: "Global",
    scopes: ["/subscriptions/187f412d-1758-44d9-b052-169e2564721d"],
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.activityLogAlerts.createOrUpdate(
    resourceGroupName,
    activityLogAlertName,
    activityLogAlertRule
  );
  console.log(result);
}

createOrUpdateAnActivityLogAlertRule().catch(console.error);
