const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to restarts one or more virtual machines in a VM scale set.
 *
 * @summary restarts one or more virtual machines in a VM scale set.
 * x-ms-original-file: 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Restart_MinimumSet_Gen.json
 */
async function virtualMachineScaleSetRestartMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  await client.virtualMachineScaleSets.restart("rgcompute", "aaaa");
}
