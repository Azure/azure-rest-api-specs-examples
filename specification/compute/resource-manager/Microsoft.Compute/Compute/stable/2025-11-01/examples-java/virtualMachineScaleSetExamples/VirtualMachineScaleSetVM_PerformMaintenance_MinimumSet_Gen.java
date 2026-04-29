
/**
 * Samples for VirtualMachineScaleSetVMs PerformMaintenance.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_PerformMaintenance_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetVM_PerformMaintenance_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachineScaleSetVMPerformMaintenanceMinimumSetGen(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSetVMs().performMaintenance("rgcompute", "aaaaaaaaaa", "aaaa",
            com.azure.core.util.Context.NONE);
    }
}
