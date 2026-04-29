
/**
 * Samples for VirtualMachineScaleSets GetOSUpgradeHistory.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_GetOSUpgradeHistory_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSet_GetOSUpgradeHistory_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachineScaleSetGetOSUpgradeHistoryMaximumSetGen(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSets().getOSUpgradeHistory("rgcompute", "aaaaaa",
            com.azure.core.util.Context.NONE);
    }
}
