const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to performs maintenance on a virtual machine in a VM scale set.
 *
 * @summary performs maintenance on a virtual machine in a VM scale set.
 * x-ms-original-file: 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_PerformMaintenance_MinimumSet_Gen.json
 */
async function virtualMachineScaleSetVMPerformMaintenanceMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  await client.virtualMachineScaleSetVMs.performMaintenance("rgcompute", "aaaaaaaaaa", "aaaa");
}
