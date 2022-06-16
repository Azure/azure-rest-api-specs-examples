const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Manual platform update domain walk to update virtual machines in a service fabric virtual machine scale set.
 *
 * @summary Manual platform update domain walk to update virtual machines in a service fabric virtual machine scale set.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_ForceRecoveryServiceFabricPlatformUpdateDomainWalk_MinimumSet_Gen.json
 */
async function virtualMachineScaleSetsForceRecoveryServiceFabricPlatformUpdateDomainWalkMinimumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const vmScaleSetName = "aaaaaaaaaaaa";
  const platformUpdateDomain = 9;
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

virtualMachineScaleSetsForceRecoveryServiceFabricPlatformUpdateDomainWalkMinimumSetGen().catch(
  console.error
);
