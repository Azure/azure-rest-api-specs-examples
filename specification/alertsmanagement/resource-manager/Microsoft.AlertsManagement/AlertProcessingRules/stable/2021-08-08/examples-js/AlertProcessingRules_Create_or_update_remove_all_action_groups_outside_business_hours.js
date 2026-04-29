const { AlertProcessingRulesManagementClient } = require("@azure/arm-alertprocessingrules");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update an alert processing rule.
 *
 * @summary create or update an alert processing rule.
 * x-ms-original-file: 2021-08-08/AlertProcessingRules_Create_or_update_remove_all_action_groups_outside_business_hours.json
 */
async function createOrUpdateARuleThatRemovesAllActionGroupsOutsideBusinessHoursMonFri09001700EasternStandardTime() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subId1";
  const client = new AlertProcessingRulesManagementClient(credential, subscriptionId);
  const result = await client.alertProcessingRules.createOrUpdate(
    "alertscorrelationrg",
    "RemoveActionGroupsOutsideBusinessHours",
    {
      location: "Global",
      properties: {
        description: "Remove all ActionGroups outside business hours",
        actions: [{ actionType: "RemoveAllActionGroups" }],
        enabled: true,
        schedule: {
          recurrences: [
            { endTime: "09:00:00", recurrenceType: "Daily", startTime: "17:00:00" },
            { daysOfWeek: ["Saturday", "Sunday"], recurrenceType: "Weekly" },
          ],
          timeZone: "Eastern Standard Time",
        },
        scopes: ["/subscriptions/subId1"],
      },
      tags: {},
    },
  );
  console.log(result);
}
