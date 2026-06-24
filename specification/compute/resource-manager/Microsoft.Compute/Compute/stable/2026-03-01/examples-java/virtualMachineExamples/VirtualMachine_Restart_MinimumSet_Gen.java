
/**
 * Samples for VirtualMachines Restart.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/virtualMachineExamples/VirtualMachine_Restart_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachine_Restart_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachineRestartMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachines().restart("rgcompute", "aaa", com.azure.core.util.Context.NONE);
    }
}
