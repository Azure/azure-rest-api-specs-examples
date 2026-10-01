const { ComputeClient } = require("@azure/arm-compute-bulkactions");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to adds resources to the specified scheduled action.
 *
 * @summary adds resources to the specified scheduled action.
 * x-ms-original-file: 2026-10-06-preview/ScheduledActions_AttachResources_PartialSuccess.json
 */
async function _03ResponseWithPartialResultsWhenAttachingResourcesToARecurringScheduledAction() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ComputeClient(credential, subscriptionId);
  const result = await client.scheduledActions.attachResources("example-rg", "weekday-start", {
    resources: [
      {
        resourceId:
          "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/web-vm-01",
      },
      {
        resourceId:
          "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/web-vm-02",
      },
    ],
  });
  console.log(result);
}
