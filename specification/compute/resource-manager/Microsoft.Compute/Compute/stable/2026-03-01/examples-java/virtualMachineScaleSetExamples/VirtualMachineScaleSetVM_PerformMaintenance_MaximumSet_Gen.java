
/**
 * Samples for VirtualMachineScaleSetVMs PerformMaintenance.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_PerformMaintenance_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetVM_PerformMaintenance_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachineScaleSetVMPerformMaintenanceMaximumSetGen(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSetVMs().performMaintenance("rgcompute", "aaaaaaaaaaaaaa",
            "aaaaaaaaaaaa", com.azure.core.util.Context.NONE);
    }
}
