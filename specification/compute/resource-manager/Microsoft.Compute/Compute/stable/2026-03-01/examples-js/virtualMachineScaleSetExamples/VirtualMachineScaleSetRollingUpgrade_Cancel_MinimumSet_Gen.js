const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to cancels the current virtual machine scale set rolling upgrade.
 *
 * @summary cancels the current virtual machine scale set rolling upgrade.
 * x-ms-original-file: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetRollingUpgrade_Cancel_MinimumSet_Gen.json
 */
async function virtualMachineScaleSetRollingUpgradeCancelMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  await client.virtualMachineScaleSetRollingUpgrades.cancel("rgcompute", "aaaaaa");
}
