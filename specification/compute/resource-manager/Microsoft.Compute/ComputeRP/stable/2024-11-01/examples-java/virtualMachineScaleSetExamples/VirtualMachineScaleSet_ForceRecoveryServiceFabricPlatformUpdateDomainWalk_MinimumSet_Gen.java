
/**
 * Samples for VirtualMachineScaleSets ForceRecoveryServiceFabricPlatformUpdateDomainWalk.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/
     * virtualMachineScaleSetExamples/
     * VirtualMachineScaleSet_ForceRecoveryServiceFabricPlatformUpdateDomainWalk_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSet_ForceRecoveryServiceFabricPlatformUpdateDomainWalk_MinimumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetForceRecoveryServiceFabricPlatformUpdateDomainWalkMinimumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSets()
            .forceRecoveryServiceFabricPlatformUpdateDomainWalkWithResponse("rgcompute", "aaaaaaaaaaaa", 9, null, null,
                com.azure.core.util.Context.NONE);
    }
}
