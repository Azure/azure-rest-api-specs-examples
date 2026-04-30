
/**
 * Samples for VirtualMachineScaleSets Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Delete_Force.json
     */
    /**
     * Sample code: Force Delete a VM scale set.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void forceDeleteAVMScaleSet(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSets().delete("myResourceGroup", "myvmScaleSet", true,
            com.azure.core.util.Context.NONE);
    }
}
