
/**
 * Samples for VirtualMachines Start.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/virtualMachineExamples/VirtualMachine_Start_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachine_Start_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachineStartMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachines().start("rgcompute", "aaaaa", com.azure.core.util.Context.NONE);
    }
}
