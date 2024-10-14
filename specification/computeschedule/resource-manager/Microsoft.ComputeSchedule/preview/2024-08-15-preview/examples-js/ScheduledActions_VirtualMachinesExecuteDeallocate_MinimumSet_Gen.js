const { ComputeScheduleClient } = require("@azure/arm-computeschedule");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to virtualMachinesExecuteDeallocate: executeDeallocate for a virtual machine
 *
 * @summary virtualMachinesExecuteDeallocate: executeDeallocate for a virtual machine
 * x-ms-original-file: 2024-08-15-preview/ScheduledActions_VirtualMachinesExecuteDeallocate_MinimumSet_Gen.json
 */
async function scheduledActionsVirtualMachinesExecuteDeallocateGeneratedByMaximumSetRuleGeneratedByMinimumSetRule() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "DE84A209-5715-43E7-BC76-3E208A9A82C5";
  const client = new ComputeScheduleClient(credential, subscriptionId);
  const result = await client.scheduledActions.virtualMachinesExecuteDeallocate(
    "irbtacrhjgfpyvhvjo",
    {
      executionParameters: {},
      resources: {
        ids: [
          "/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Compute/virtualMachines/testResource2",
        ],
      },
      correlationid: "01080d2f-1dca-4610-afb4-dd25eddf1f3c1",
    },
  );
  console.log(result);
}
