
/**
 * Samples for VirtualMachineScaleSets ForceRecoveryServiceFabricPlatformUpdateDomainWalk.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/virtualMachineScaleSetExamples/
     * VirtualMachineScaleSet_ForceRecoveryServiceFabricPlatformUpdateDomainWalk_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSet_ForceRecoveryServiceFabricPlatformUpdateDomainWalk_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachineScaleSetForceRecoveryServiceFabricPlatformUpdateDomainWalkMaximumSetGen(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSets()
            .forceRecoveryServiceFabricPlatformUpdateDomainWalkWithResponse("rgcompute", "aaaaaaaaaaaaaaaa", 30, null,
                null, com.azure.core.util.Context.NONE);
    }
}
