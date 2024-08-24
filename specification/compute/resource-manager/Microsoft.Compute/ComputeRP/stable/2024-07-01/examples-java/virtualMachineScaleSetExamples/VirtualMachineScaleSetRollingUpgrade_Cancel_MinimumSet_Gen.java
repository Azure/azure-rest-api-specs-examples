
/**
 * Samples for VirtualMachineScaleSetRollingUpgrades Cancel.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/
     * virtualMachineScaleSetExamples/VirtualMachineScaleSetRollingUpgrade_Cancel_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetRollingUpgrade_Cancel_MinimumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        virtualMachineScaleSetRollingUpgradeCancelMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSetRollingUpgrades().cancel("rgcompute",
            "aaaaaa", com.azure.core.util.Context.NONE);
    }
}
