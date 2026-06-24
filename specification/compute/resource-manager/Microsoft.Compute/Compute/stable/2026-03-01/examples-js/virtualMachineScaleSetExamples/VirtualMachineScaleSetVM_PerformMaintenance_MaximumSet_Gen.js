const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to performs maintenance on a virtual machine in a VM scale set.
 *
 * @summary performs maintenance on a virtual machine in a VM scale set.
 * x-ms-original-file: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_PerformMaintenance_MaximumSet_Gen.json
 */
async function virtualMachineScaleSetVMPerformMaintenanceMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  await client.virtualMachineScaleSetVMs.performMaintenance(
    "rgcompute",
    "aaaaaaaaaaaaaa",
    "aaaaaaaaaaaa",
  );
}
