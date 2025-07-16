const { ComputeScheduleClient } = require("@azure/arm-computeschedule");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to virtualMachinesSubmitStart: Schedule start operation for a batch of virtual machines at datetime in future.
 *
 * @summary virtualMachinesSubmitStart: Schedule start operation for a batch of virtual machines at datetime in future.
 * x-ms-original-file: 2025-05-01/ScheduledActions_VirtualMachinesSubmitStart_MinimumSet_Gen.json
 */
async function scheduledActionsVirtualMachinesSubmitStartMinimumSetGenGeneratedByMinimumSetRule() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "0505D8E4-D41A-48FB-9CA5-4AF8D93BE75F";
  const client = new ComputeScheduleClient(credential, subscriptionId);
  const result = await client.scheduledActions.virtualMachinesSubmitStart("hdttrxnbswit", {
    schedule: { deadlineType: "Unknown" },
    executionParameters: {},
    resources: {
      ids: [
        "/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Compute/virtualMachines/testResource3",
      ],
    },
    correlationId: "bvmpxvbd",
  });
  console.log(result);
}
