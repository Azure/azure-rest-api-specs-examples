
/**
 * Samples for VirtualMachineScaleSetRollingUpgrades Cancel.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetRollingUpgrade_Cancel_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetRollingUpgrade_Cancel_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachineScaleSetRollingUpgradeCancelMaximumSetGen(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSetRollingUpgrades().cancel("rgcompute", "aaaaa",
            com.azure.core.util.Context.NONE);
    }
}
