const { ComputeClient } = require("@azure/arm-compute-bulkactions");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates the specified scheduled action.
 *
 * @summary updates the specified scheduled action.
 * x-ms-original-file: 2026-10-06-preview/ScheduledActions_Update_BasicSuccess.json
 */
async function _01UpdateTheActionTypeOfARecurringScheduledAction() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ComputeClient(credential, subscriptionId);
  await client.scheduledActions.update("example-rg", "weekday-start", {
    properties: { actionType: "Deallocate" },
  });
}
