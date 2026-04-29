
/**
 * Samples for VirtualMachines Restart.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/virtualMachineExamples/VirtualMachine_Restart_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachine_Restart_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachineRestartMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachines().restart("rgcompute", "aaaaaaaaaaaaaaaaaaaa",
            com.azure.core.util.Context.NONE);
    }
}
