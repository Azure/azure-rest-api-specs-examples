import com.azure.core.util.Context;

/** Samples for VirtualMachineScaleSets ForceRecoveryServiceFabricPlatformUpdateDomainWalk. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_ForceRecoveryServiceFabricPlatformUpdateDomainWalk_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSets_ForceRecoveryServiceFabricPlatformUpdateDomainWalk_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetsForceRecoveryServiceFabricPlatformUpdateDomainWalkMaximumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSets()
            .forceRecoveryServiceFabricPlatformUpdateDomainWalkWithResponse(
                "rgcompute", "aaaaaaaaaaaaaaaa", 30, null, null, Context.NONE);
    }
}
