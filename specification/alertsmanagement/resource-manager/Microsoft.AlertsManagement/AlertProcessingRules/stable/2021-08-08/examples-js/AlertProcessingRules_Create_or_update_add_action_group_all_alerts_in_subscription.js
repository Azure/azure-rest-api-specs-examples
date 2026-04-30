const { AlertProcessingRulesManagementClient } = require("@azure/arm-alertprocessingrules");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update an alert processing rule.
 *
 * @summary create or update an alert processing rule.
 * x-ms-original-file: 2021-08-08/AlertProcessingRules_Create_or_update_add_action_group_all_alerts_in_subscription.json
 */
async function createOrUpdateARuleThatAddsAnActionGroupToAllAlertsInASubscription() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subId1";
  const client = new AlertProcessingRulesManagementClient(credential, subscriptionId);
  const result = await client.alertProcessingRules.createOrUpdate(
    "alertscorrelationrg",
    "AddActionGroupToSubscription",
    {
      location: "Global",
      properties: {
        description: "Add ActionGroup1 to all alerts in the subscription",
        actions: [
          {
            actionGroupIds: [
              "/subscriptions/subId1/resourcegroups/RGId1/providers/microsoft.insights/actiongroups/ActionGroup1",
            ],
            actionType: "AddActionGroups",
          },
        ],
        enabled: true,
        scopes: ["/subscriptions/subId1"],
      },
      tags: {},
    },
  );
  console.log(result);
}
