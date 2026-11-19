const { CognitiveServicesManagementClient } = require("@azure/arm-cognitiveservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates a Foundry project cost control. Omitting properties preserves all existing cost control properties; setting properties to null or an empty object ({}) clears them all.
 * A non-empty properties object updates selected fields, preserving omitted fields; a successful update that leaves the cost control unchanged retains its entity tag.
 * The complete merged result may contain at most one notification, including disabled notifications, and one Block action.
 * Exceeding either limit returns HTTP 400 without changing the stored definition.
 *
 * @summary updates a Foundry project cost control. Omitting properties preserves all existing cost control properties; setting properties to null or an empty object ({}) clears them all.
 * A non-empty properties object updates selected fields, preserving omitted fields; a successful update that leaves the cost control unchanged retains its entity tag.
 * The complete merged result may contain at most one notification, including disabled notifications, and one Block action.
 * Exceeding either limit returns HTTP 400 without changing the stored definition.
 * x-ms-original-file: 2026-11-15-preview/ProjectCostControl/clearOptionalSettings.json
 */
async function clearFiltersAndRemoveNamedNotificationsAndBlockingActions() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new CognitiveServicesManagementClient(credential, subscriptionId);
  const result = await client.projectCostControls.update(
    "foundry-resource-group",
    "foundry-account",
    "production",
    "project-token-budget",
    {
      properties: {
        filter: {},
        notifications: { usageAlert: {} },
        providerProperties: { actions: { stopRequests: {} } },
      },
    },
    { ifMatch: '"00000000-0000-0000-0000-000000000001"' },
  );
  console.log(result);
}
