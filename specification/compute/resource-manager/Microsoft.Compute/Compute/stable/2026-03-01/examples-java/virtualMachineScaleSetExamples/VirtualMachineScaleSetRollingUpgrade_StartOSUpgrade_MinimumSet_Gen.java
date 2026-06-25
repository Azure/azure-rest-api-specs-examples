
/**
 * Samples for VirtualMachineScaleSetRollingUpgrades StartOSUpgrade.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetRollingUpgrade_StartOSUpgrade_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetRollingUpgrade_StartOSUpgrade_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachineScaleSetRollingUpgradeStartOSUpgradeMinimumSetGen(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSetRollingUpgrades().startOSUpgrade("rgcompute",
            "aaaaaaaaaaaaaaaaaa", com.azure.core.util.Context.NONE);
    }
}
