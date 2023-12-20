const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Manual platform update domain walk to update virtual machines in a service fabric virtual machine scale set.
 *
 * @summary Manual platform update domain walk to update virtual machines in a service fabric virtual machine scale set.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSet_ForceRecoveryServiceFabricPlatformUpdateDomainWalk_MaximumSet_Gen.json
 */
async function virtualMachineScaleSetForceRecoveryServiceFabricPlatformUpdateDomainWalkMaximumSetGen() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "rgcompute";
  const vmScaleSetName = "aaaaaaaaaaaaaaaa";
  const platformUpdateDomain = 30;
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result =
    await client.virtualMachineScaleSets.forceRecoveryServiceFabricPlatformUpdateDomainWalk(
      resourceGroupName,
      vmScaleSetName,
      platformUpdateDomain,
    );
  console.log(result);
}
