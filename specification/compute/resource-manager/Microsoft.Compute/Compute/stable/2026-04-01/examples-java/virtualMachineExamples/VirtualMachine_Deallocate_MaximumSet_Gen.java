
/**
 * Samples for VirtualMachines Deallocate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/virtualMachineExamples/VirtualMachine_Deallocate_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachine_Deallocate_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachineDeallocateMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachines().deallocate("rgcompute", "aaaaaaaaaa", true, null,
            com.azure.core.util.Context.NONE);
    }
}
