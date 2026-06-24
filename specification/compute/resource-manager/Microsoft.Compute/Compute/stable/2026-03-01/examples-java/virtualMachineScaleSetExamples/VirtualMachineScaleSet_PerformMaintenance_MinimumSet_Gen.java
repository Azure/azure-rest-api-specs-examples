
/**
 * Samples for VirtualMachineScaleSets PerformMaintenance.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_PerformMaintenance_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSet_PerformMaintenance_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachineScaleSetPerformMaintenanceMinimumSetGen(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSets().performMaintenance("rgcompute", "aa", null,
            com.azure.core.util.Context.NONE);
    }
}
