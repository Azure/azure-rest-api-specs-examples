const { ComputeScheduleClient } = require("@azure/arm-computeschedule");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to virtualMachinesSubmitDeallocate: Schedule deallocate operation for a batch of virtual machines at datetime in future.
 *
 * @summary virtualMachinesSubmitDeallocate: Schedule deallocate operation for a batch of virtual machines at datetime in future.
 * x-ms-original-file: 2025-05-01/ScheduledActions_VirtualMachinesSubmitDeallocate_MinimumSet_Gen.json
 */
async function scheduledActionsVirtualMachinesSubmitDeallocateMinimumSetGenGeneratedByMinimumSetRule() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "0505D8E4-D41A-48FB-9CA5-4AF8D93BE75F";
  const client = new ComputeScheduleClient(credential, subscriptionId);
  const result = await client.scheduledActions.virtualMachinesSubmitDeallocate(
    "ccrsyfkiakaxblrddurmxbju",
    {
      schedule: { deadlineType: "Unknown" },
      executionParameters: {},
      resources: {
        ids: [
          "/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Compute/virtualMachines/testResource3",
        ],
      },
      correlationId: "evmwonebfzxenjdpucgcwdjdya",
    },
  );
  console.log(result);
}
