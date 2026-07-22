
/**
 * Samples for VirtualMachines Deallocate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/virtualMachineExamples/VirtualMachine_Deallocate_WithForceDeallocate.json
     */
    /**
     * Sample code: VirtualMachine_Deallocate_WithForceDeallocate.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineDeallocateWithForceDeallocate(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachines().deallocate("rgcompute", "aaaaaaaaaaaaaaaa", null, true,
            com.azure.core.util.Context.NONE);
    }
}
