const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets the status of the latest virtual machine scale set rolling upgrade.
 *
 * @summary gets the status of the latest virtual machine scale set rolling upgrade.
 * x-ms-original-file: 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetRollingUpgrade_GetLatest_MaximumSet_Gen.json
 */
async function virtualMachineScaleSetRollingUpgradeGetLatestMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSetRollingUpgrades.getLatest(
    "rgcompute",
    "aaaaaaaaaaaaaaaaaaaaaaaaa",
  );
  console.log(result);
}
