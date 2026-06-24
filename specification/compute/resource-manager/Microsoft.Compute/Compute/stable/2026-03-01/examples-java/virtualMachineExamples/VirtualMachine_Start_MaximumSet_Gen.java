
/**
 * Samples for VirtualMachines Start.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/virtualMachineExamples/VirtualMachine_Start_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachine_Start_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachineStartMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachines().start("rgcompute", "aaaaaaaaaaaaaaaaaaaa",
            com.azure.core.util.Context.NONE);
    }
}
