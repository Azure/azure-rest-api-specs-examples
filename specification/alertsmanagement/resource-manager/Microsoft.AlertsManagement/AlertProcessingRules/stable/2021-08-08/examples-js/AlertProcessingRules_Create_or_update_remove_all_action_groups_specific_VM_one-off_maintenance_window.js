const { AlertProcessingRulesManagementClient } = require("@azure/arm-alertprocessingrules");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update an alert processing rule.
 *
 * @summary create or update an alert processing rule.
 * x-ms-original-file: 2021-08-08/AlertProcessingRules_Create_or_update_remove_all_action_groups_specific_VM_one-off_maintenance_window.json
 */
async function createOrUpdateARuleThatRemovesAllActionGroupsFromAlertsOnASpecificVMDuringAOneOffMaintenanceWindow18002000AtASpecificDatePacificStandardTime() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subId1";
  const client = new AlertProcessingRulesManagementClient(credential, subscriptionId);
  const result = await client.alertProcessingRules.createOrUpdate(
    "alertscorrelationrg",
    "RemoveActionGroupsMaintenanceWindow",
    {
      location: "Global",
      properties: {
        description:
          "Removes all ActionGroups from all Alerts on VMName during the maintenance window",
        actions: [{ actionType: "RemoveAllActionGroups" }],
        enabled: true,
        schedule: {
          effectiveFrom: "2021-04-15T18:00:00",
          effectiveUntil: "2021-04-15T20:00:00",
          timeZone: "Pacific Standard Time",
        },
        scopes: [
          "/subscriptions/subId1/resourceGroups/RGId1/providers/Microsoft.Compute/virtualMachines/VMName",
        ],
      },
      tags: {},
    },
  );
  console.log(result);
}
