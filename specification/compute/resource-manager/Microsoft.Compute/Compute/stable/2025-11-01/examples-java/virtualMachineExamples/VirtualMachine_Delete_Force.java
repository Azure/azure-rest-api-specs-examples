
/**
 * Samples for VirtualMachines Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/virtualMachineExamples/VirtualMachine_Delete_Force.json
     */
    /**
     * Sample code: Force delete a VM.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void forceDeleteAVM(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachines().delete("myResourceGroup", "myVM", true,
            com.azure.core.util.Context.NONE);
    }
}
