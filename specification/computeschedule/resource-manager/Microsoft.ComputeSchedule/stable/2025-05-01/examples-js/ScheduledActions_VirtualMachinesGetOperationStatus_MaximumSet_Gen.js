const { ComputeScheduleClient } = require("@azure/arm-computeschedule");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to virtualMachinesGetOperationStatus: Polling endpoint to read status of operations performed on virtual machines
 *
 * @summary virtualMachinesGetOperationStatus: Polling endpoint to read status of operations performed on virtual machines
 * x-ms-original-file: 2025-05-01/ScheduledActions_VirtualMachinesGetOperationStatus_MaximumSet_Gen.json
 */
async function scheduledActionsVirtualMachinesGetOperationStatusMaximumSetGenGeneratedByMaximumSetRule() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "0505D8E4-D41A-48FB-9CA5-4AF8D93BE75F";
  const client = new ComputeScheduleClient(credential, subscriptionId);
  const result = await client.scheduledActions.virtualMachinesGetOperationStatus(
    "fvyydtrzewnjtrezzieqxs",
    {
      operationIds: ["hswzfrierpxdgcuu"],
      correlationId: "jtlszorevrftvfhnqoxlwpiwcbmj",
    },
  );
  console.log(result);
}
