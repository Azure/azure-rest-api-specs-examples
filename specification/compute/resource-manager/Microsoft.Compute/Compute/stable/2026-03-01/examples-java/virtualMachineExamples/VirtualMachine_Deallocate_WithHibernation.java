
/**
 * Samples for VirtualMachines Deallocate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/virtualMachineExamples/VirtualMachine_Deallocate_WithHibernation.json
     */
    /**
     * Sample code: VirtualMachine_Deallocate_WithHibernation.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineDeallocateWithHibernation(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachines().deallocate("rgcompute", "aaaaaaaaaaaaaaaa", true, null,
            com.azure.core.util.Context.NONE);
    }
}
