const { AlertProcessingRulesManagementClient } = require("@azure/arm-alertprocessingrules");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update an alert processing rule.
 *
 * @summary create or update an alert processing rule.
 * x-ms-original-file: 2021-08-08/AlertProcessingRules_Create_or_update_remove_all_action_groups_recurring_maintenance_window.json
 */
async function createOrUpdateARuleThatRemovesAllActionGroupsFromAllAlertsOnAnyVMInTwoResourceGroupsDuringARecurringMaintenanceWindow22000400EverySatAndSunIndiaStandardTime() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subId1";
  const client = new AlertProcessingRulesManagementClient(credential, subscriptionId);
  const result = await client.alertProcessingRules.createOrUpdate(
    "alertscorrelationrg",
    "RemoveActionGroupsRecurringMaintenance",
    {
      location: "Global",
      properties: {
        description:
          "Remove all ActionGroups from all Vitual machine Alerts during the recurring maintenance",
        actions: [{ actionType: "RemoveAllActionGroups" }],
        conditions: [
          {
            field: "TargetResourceType",
            operator: "Equals",
            values: ["microsoft.compute/virtualmachines"],
          },
        ],
        enabled: true,
        schedule: {
          recurrences: [
            {
              daysOfWeek: ["Saturday", "Sunday"],
              endTime: "04:00:00",
              recurrenceType: "Weekly",
              startTime: "22:00:00",
            },
          ],
          timeZone: "India Standard Time",
        },
        scopes: [
          "/subscriptions/subId1/resourceGroups/RGId1",
          "/subscriptions/subId1/resourceGroups/RGId2",
        ],
      },
      tags: {},
    },
  );
  console.log(result);
}
