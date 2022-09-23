const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Manual platform update domain walk to update virtual machines in a service fabric virtual machine scale set.
 *
 * @summary Manual platform update domain walk to update virtual machines in a service fabric virtual machine scale set.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSets_ForceRecoveryServiceFabricPlatformUpdateDomainWalk_MaximumSet_Gen.json
 */
async function virtualMachineScaleSetsForceRecoveryServiceFabricPlatformUpdateDomainWalkMaximumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const vmScaleSetName = "aaaaaaaaaaaaaaaa";
  const platformUpdateDomain = 30;
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result =
    await client.virtualMachineScaleSets.forceRecoveryServiceFabricPlatformUpdateDomainWalk(
      resourceGroupName,
      vmScaleSetName,
      platformUpdateDomain
    );
  console.log(result);
}

virtualMachineScaleSetsForceRecoveryServiceFabricPlatformUpdateDomainWalkMaximumSetGen().catch(
  console.error
);
