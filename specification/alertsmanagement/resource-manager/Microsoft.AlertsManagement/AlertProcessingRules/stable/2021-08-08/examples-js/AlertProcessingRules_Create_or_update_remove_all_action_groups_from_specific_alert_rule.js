const { AlertProcessingRulesManagementClient } = require("@azure/arm-alertprocessingrules");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update an alert processing rule.
 *
 * @summary create or update an alert processing rule.
 * x-ms-original-file: 2021-08-08/AlertProcessingRules_Create_or_update_remove_all_action_groups_from_specific_alert_rule.json
 */
async function createOrUpdateARuleThatRemovesAllActionGroupsFromAllAlertsInASubscriptionComingFromASpecificAlertRule() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subId1";
  const client = new AlertProcessingRulesManagementClient(credential, subscriptionId);
  const result = await client.alertProcessingRules.createOrUpdate(
    "alertscorrelationrg",
    "RemoveActionGroupsSpecificAlertRule",
    {
      location: "Global",
      properties: {
        description: "Removes all ActionGroups from all Alerts that fire on above AlertRule",
        actions: [{ actionType: "RemoveAllActionGroups" }],
        conditions: [
          {
            field: "AlertRuleId",
            operator: "Equals",
            values: [
              "/subscriptions/suubId1/resourceGroups/Rgid2/providers/microsoft.insights/activityLogAlerts/RuleName",
            ],
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
