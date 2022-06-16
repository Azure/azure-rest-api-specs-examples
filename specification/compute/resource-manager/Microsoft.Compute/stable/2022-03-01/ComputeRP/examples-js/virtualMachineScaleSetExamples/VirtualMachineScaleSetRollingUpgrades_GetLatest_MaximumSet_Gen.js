const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the status of the latest virtual machine scale set rolling upgrade.
 *
 * @summary Gets the status of the latest virtual machine scale set rolling upgrade.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetRollingUpgrades_GetLatest_MaximumSet_Gen.json
 */
async function virtualMachineScaleSetRollingUpgradesGetLatestMaximumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const vmScaleSetName = "aaaaaaaaaaaaaaaaaaaaaaaaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSetRollingUpgrades.getLatest(
    resourceGroupName,
    vmScaleSetName
  );
  console.log(result);
}

virtualMachineScaleSetRollingUpgradesGetLatestMaximumSetGen().catch(console.error);
