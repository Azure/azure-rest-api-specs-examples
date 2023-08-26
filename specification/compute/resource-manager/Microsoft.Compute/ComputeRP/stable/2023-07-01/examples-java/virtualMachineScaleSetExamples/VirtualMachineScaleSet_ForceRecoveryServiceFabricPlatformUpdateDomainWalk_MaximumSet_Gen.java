/** Samples for VirtualMachineScaleSets ForceRecoveryServiceFabricPlatformUpdateDomainWalk. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSet_ForceRecoveryServiceFabricPlatformUpdateDomainWalk_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSet_ForceRecoveryServiceFabricPlatformUpdateDomainWalk_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetForceRecoveryServiceFabricPlatformUpdateDomainWalkMaximumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSets()
            .forceRecoveryServiceFabricPlatformUpdateDomainWalkWithResponse(
                "rgcompute", "aaaaaaaaaaaaaaaa", 30, null, null, com.azure.core.util.Context.NONE);
    }
}
