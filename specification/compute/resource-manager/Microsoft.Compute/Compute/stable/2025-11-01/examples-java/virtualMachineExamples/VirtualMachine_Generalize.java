
/**
 * Samples for VirtualMachines Generalize.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/virtualMachineExamples/VirtualMachine_Generalize.json
     */
    /**
     * Sample code: Generalize a Virtual Machine.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void generalizeAVirtualMachine(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachines().generalizeWithResponse("myResourceGroup", "myVMName",
            com.azure.core.util.Context.NONE);
    }
}
