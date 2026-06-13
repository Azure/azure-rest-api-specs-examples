const { ComputeScheduleClient } = require("@azure/arm-computeschedule");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to [PRIVATE PREVIEW]: VirtualMachinesExecuteCreate: Execute create operation for a batch of virtual machines, this operation is triggered as soon as Computeschedule receives it.
 *
 * @summary [PRIVATE PREVIEW]: VirtualMachinesExecuteCreate: Execute create operation for a batch of virtual machines, this operation is triggered as soon as Computeschedule receives it.
 * x-ms-original-file: 2026-04-15-preview/ScheduledActions_VirtualMachinesExecuteCreate_MinimumSet_Gen.json
 */
async function scheduledActionsVirtualMachinesExecuteCreateMinimumSet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "732116BD-AF31-4E74-9283-B387C44B4A44";
  const client = new ComputeScheduleClient(credential, subscriptionId);
  const result = await client.scheduledActions.virtualMachinesExecuteCreate("eastus2", {
    resourceConfigParameters: { resourceCount: 3 },
    executionParameters: {},
  });
  console.log(result);
}
