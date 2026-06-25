
/**
 * Samples for VirtualMachines PerformMaintenance.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/virtualMachineExamples/VirtualMachine_PerformMaintenance_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachine_PerformMaintenance_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachinePerformMaintenanceMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachines().performMaintenance("rgcompute", "aaaaaaaaaa",
            com.azure.core.util.Context.NONE);
    }
}
