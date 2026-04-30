const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to restart a virtual machine.
 *
 * @summary the operation to restart a virtual machine.
 * x-ms-original-file: 2025-11-01/virtualMachineExamples/VirtualMachine_Restart_MaximumSet_Gen.json
 */
async function virtualMachineRestartMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  await client.virtualMachines.restart("rgcompute", "aaaaaaaaaaaaaaaaaaaa");
}
