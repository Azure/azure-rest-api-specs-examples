const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Cancels the current virtual machine scale set rolling upgrade.
 *
 * @summary Cancels the current virtual machine scale set rolling upgrade.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-03-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetRollingUpgrade_Cancel_MinimumSet_Gen.json
 */
async function virtualMachineScaleSetRollingUpgradeCancelMinimumSetGen() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "rgcompute";
  const vmScaleSetName = "aaaaaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSetRollingUpgrades.beginCancelAndWait(
    resourceGroupName,
    vmScaleSetName
  );
  console.log(result);
}
