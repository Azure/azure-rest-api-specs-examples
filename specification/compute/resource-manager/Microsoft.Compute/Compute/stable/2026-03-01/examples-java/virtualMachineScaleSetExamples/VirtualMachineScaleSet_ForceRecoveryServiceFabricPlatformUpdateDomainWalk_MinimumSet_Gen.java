
/**
 * Samples for VirtualMachineScaleSets ForceRecoveryServiceFabricPlatformUpdateDomainWalk.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/virtualMachineScaleSetExamples/
     * VirtualMachineScaleSet_ForceRecoveryServiceFabricPlatformUpdateDomainWalk_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSet_ForceRecoveryServiceFabricPlatformUpdateDomainWalk_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachineScaleSetForceRecoveryServiceFabricPlatformUpdateDomainWalkMinimumSetGen(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSets()
            .forceRecoveryServiceFabricPlatformUpdateDomainWalkWithResponse("rgcompute", "aaaaaaaaaaaa", 9, null, null,
                com.azure.core.util.Context.NONE);
    }
}
