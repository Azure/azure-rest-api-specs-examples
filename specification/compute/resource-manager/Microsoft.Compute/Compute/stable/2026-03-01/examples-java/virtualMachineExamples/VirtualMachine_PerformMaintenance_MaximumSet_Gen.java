
/**
 * Samples for VirtualMachines PerformMaintenance.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/virtualMachineExamples/VirtualMachine_PerformMaintenance_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachine_PerformMaintenance_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachinePerformMaintenanceMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachines().performMaintenance("rgcompute", "aaaaaaa",
            com.azure.core.util.Context.NONE);
    }
}
