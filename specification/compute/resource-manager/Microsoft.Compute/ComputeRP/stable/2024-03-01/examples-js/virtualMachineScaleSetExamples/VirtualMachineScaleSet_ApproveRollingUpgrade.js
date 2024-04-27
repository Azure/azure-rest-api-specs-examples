const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Approve upgrade on deferred rolling upgrades for OS disks in the virtual machines in a VM scale set.
 *
 * @summary Approve upgrade on deferred rolling upgrades for OS disks in the virtual machines in a VM scale set.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSet_ApproveRollingUpgrade.json
 */
async function virtualMachineScaleSetApproveRollingUpgrade() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "rgcompute";
  const vmScaleSetName = "vmssToApproveRollingUpgradeOn";
  const vmInstanceIDs = {
    instanceIds: ["0", "1", "2"],
  };
  const options = {
    vmInstanceIDs,
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSets.beginApproveRollingUpgradeAndWait(
    resourceGroupName,
    vmScaleSetName,
    options,
  );
  console.log(result);
}
