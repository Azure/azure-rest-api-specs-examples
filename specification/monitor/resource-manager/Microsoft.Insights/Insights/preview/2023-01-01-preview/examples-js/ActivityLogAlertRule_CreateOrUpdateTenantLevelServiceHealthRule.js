const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a new Activity Log Alert rule or update an existing one.
 *
 * @summary create a new Activity Log Alert rule or update an existing one.
 * x-ms-original-file: 2023-01-01-preview/ActivityLogAlertRule_CreateOrUpdateTenantLevelServiceHealthRule.json
 */
async function createOrUpdateAnActivityLogAlertRuleForTenantLevelEvents() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "187f412d-1758-44d9-b052-169e2564721d";
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.activityLogAlerts.createOrUpdate(
    "MyResourceGroup",
    "SampleActivityLogAlertSHRuleOnTenantLevel",
    {
      location: "Global",
      description:
        "Description of sample Activity Log Alert service health rule on tenant level events.",
      actions: {
        actionGroups: [
          {
            actionGroupId:
              "/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/MyResourceGroup/providers/Microsoft.Insights/actionGroups/SampleActionGroup",
            actionProperties: { "Email.Title": "my email title" },
            webhookProperties: { sampleWebhookProperty: "SamplePropertyValue" },
          },
        ],
      },
      condition: { allOf: [{ equals: "ServiceHealth", field: "category" }] },
      enabled: true,
      tenantScope: "72f988bf-86f1-41af-91ab-2d7cd011db47",
      tags: {},
    },
  );
  console.log(result);
}
