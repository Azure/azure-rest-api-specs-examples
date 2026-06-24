
/**
 * Samples for VirtualMachines Reapply.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/virtualMachineExamples/VirtualMachine_Reapply.json
     */
    /**
     * Sample code: Reapply the state of a virtual machine.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void reapplyTheStateOfAVirtualMachine(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachines().reapply("ResourceGroup", "VMName",
            com.azure.core.util.Context.NONE);
    }
}
