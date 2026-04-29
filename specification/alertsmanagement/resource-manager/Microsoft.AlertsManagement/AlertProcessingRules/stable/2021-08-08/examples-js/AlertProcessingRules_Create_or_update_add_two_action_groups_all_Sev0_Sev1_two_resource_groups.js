const { AlertProcessingRulesManagementClient } = require("@azure/arm-alertprocessingrules");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update an alert processing rule.
 *
 * @summary create or update an alert processing rule.
 * x-ms-original-file: 2021-08-08/AlertProcessingRules_Create_or_update_add_two_action_groups_all_Sev0_Sev1_two_resource_groups.json
 */
async function createOrUpdateARuleThatAddsTwoActionGroupsToAllSev0AndSev1AlertsInTwoResourceGroups() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subId1";
  const client = new AlertProcessingRulesManagementClient(credential, subscriptionId);
  const result = await client.alertProcessingRules.createOrUpdate(
    "alertscorrelationrg",
    "AddActionGroupsBySeverity",
    {
      location: "Global",
      properties: {
        description: "Add AGId1 and AGId2 to all Sev0 and Sev1 alerts in these resourceGroups",
        actions: [
          {
            actionGroupIds: [
              "/subscriptions/subId1/resourcegroups/RGId1/providers/microsoft.insights/actiongroups/AGId1",
              "/subscriptions/subId1/resourcegroups/RGId1/providers/microsoft.insights/actiongroups/AGId2",
            ],
            actionType: "AddActionGroups",
          },
        ],
        conditions: [{ field: "Severity", operator: "Equals", values: ["sev0", "sev1"] }],
        enabled: true,
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
