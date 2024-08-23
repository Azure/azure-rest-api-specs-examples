
/**
 * Samples for VirtualMachineScaleSetRollingUpgrades StartOSUpgrade.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/
     * virtualMachineScaleSetExamples/VirtualMachineScaleSetRollingUpgrade_StartOSUpgrade_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetRollingUpgrade_StartOSUpgrade_MinimumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetRollingUpgradeStartOSUpgradeMinimumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSetRollingUpgrades()
            .startOSUpgrade("rgcompute", "aaaaaaaaaaaaaaaaaa", com.azure.core.util.Context.NONE);
    }
}
