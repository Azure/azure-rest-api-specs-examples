
/**
 * Samples for VirtualMachineScaleSets Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/
     * virtualMachineScaleSetExamples/VirtualMachineScaleSet_Delete_Force.json
     */
    /**
     * Sample code: Force Delete a VM scale set.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void forceDeleteAVMScaleSet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSets().delete("myResourceGroup",
            "myvmScaleSet", true, com.azure.core.util.Context.NONE);
    }
}
