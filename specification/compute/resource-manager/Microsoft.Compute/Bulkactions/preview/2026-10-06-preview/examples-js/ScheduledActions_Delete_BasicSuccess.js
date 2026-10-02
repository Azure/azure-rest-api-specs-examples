const { ComputeClient } = require("@azure/arm-compute-bulkactions");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to deletes the specified scheduled action.
 *
 * @summary deletes the specified scheduled action.
 * x-ms-original-file: 2026-10-06-preview/ScheduledActions_Delete_BasicSuccess.json
 */
async function deleteARecurringScheduledAction() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ComputeClient(credential, subscriptionId);
  await client.scheduledActions.delete("example-rg", "weekday-start");
}
