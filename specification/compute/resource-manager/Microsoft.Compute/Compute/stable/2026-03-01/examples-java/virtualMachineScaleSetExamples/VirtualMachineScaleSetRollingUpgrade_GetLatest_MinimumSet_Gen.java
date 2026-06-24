
/**
 * Samples for VirtualMachineScaleSetRollingUpgrades GetLatest.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetRollingUpgrade_GetLatest_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetRollingUpgrade_GetLatest_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachineScaleSetRollingUpgradeGetLatestMinimumSetGen(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSetRollingUpgrades().getLatestWithResponse("rgcompute",
            "aaaaaaaaaaaaaaaaa", com.azure.core.util.Context.NONE);
    }
}
