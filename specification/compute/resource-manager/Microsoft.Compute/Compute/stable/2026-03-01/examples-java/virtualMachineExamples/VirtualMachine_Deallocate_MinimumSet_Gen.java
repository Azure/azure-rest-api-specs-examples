
/**
 * Samples for VirtualMachines Deallocate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/virtualMachineExamples/VirtualMachine_Deallocate_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachine_Deallocate_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachineDeallocateMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachines().deallocate("rgcompute", "aaaaaaaaaaaaaaaa", null, null,
            com.azure.core.util.Context.NONE);
    }
}
