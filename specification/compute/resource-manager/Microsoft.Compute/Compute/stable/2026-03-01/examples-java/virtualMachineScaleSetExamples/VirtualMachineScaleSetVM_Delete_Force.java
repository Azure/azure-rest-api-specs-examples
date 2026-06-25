
/**
 * Samples for VirtualMachineScaleSetVMs Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_Delete_Force.json
     */
    /**
     * Sample code: Force Delete a virtual machine from a VM scale set.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        forceDeleteAVirtualMachineFromAVMScaleSet(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSetVMs().delete("myResourceGroup", "myvmScaleSet", "0", true,
            com.azure.core.util.Context.NONE);
    }
}
