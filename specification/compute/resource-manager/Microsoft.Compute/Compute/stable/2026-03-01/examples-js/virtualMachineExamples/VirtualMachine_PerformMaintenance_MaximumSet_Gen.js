const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to perform maintenance on a virtual machine.
 *
 * @summary the operation to perform maintenance on a virtual machine.
 * x-ms-original-file: 2026-03-01/virtualMachineExamples/VirtualMachine_PerformMaintenance_MaximumSet_Gen.json
 */
async function virtualMachinePerformMaintenanceMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  await client.virtualMachines.performMaintenance("rgcompute", "aaaaaaa");
}
