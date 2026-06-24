const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets the status of the latest virtual machine scale set rolling upgrade.
 *
 * @summary gets the status of the latest virtual machine scale set rolling upgrade.
 * x-ms-original-file: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetRollingUpgrade_GetLatest_MinimumSet_Gen.json
 */
async function virtualMachineScaleSetRollingUpgradeGetLatestMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSetRollingUpgrades.getLatest(
    "rgcompute",
    "aaaaaaaaaaaaaaaaa",
  );
  console.log(result);
}
