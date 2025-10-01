
/**
 * Samples for VirtualMachineScaleSetRollingUpgrades StartOSUpgrade.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/
     * virtualMachineScaleSetExamples/VirtualMachineScaleSetRollingUpgrade_StartOSUpgrade_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetRollingUpgrade_StartOSUpgrade_MaximumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetRollingUpgradeStartOSUpgradeMaximumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSetRollingUpgrades()
            .startOSUpgrade("rgcompute", "aaaa", com.azure.core.util.Context.NONE);
    }
}
