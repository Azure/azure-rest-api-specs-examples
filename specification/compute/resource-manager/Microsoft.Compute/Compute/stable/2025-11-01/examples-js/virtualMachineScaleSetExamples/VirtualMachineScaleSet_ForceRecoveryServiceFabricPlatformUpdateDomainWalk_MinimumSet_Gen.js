const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to manual platform update domain walk to update virtual machines in a service fabric virtual machine scale set.
 *
 * @summary manual platform update domain walk to update virtual machines in a service fabric virtual machine scale set.
 * x-ms-original-file: 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_ForceRecoveryServiceFabricPlatformUpdateDomainWalk_MinimumSet_Gen.json
 */
async function virtualMachineScaleSetForceRecoveryServiceFabricPlatformUpdateDomainWalkMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result =
    await client.virtualMachineScaleSets.forceRecoveryServiceFabricPlatformUpdateDomainWalk(
      "rgcompute",
      "aaaaaaaaaaaa",
      9,
    );
  console.log(result);
}
