
/**
 * Samples for VirtualMachineScaleSetVMs Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/
     * virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_Delete_Force.json
     */
    /**
     * Sample code: Force Delete a virtual machine from a VM scale set.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void forceDeleteAVirtualMachineFromAVMScaleSet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSetVMs().delete("myResourceGroup",
            "myvmScaleSet", "0", true, com.azure.core.util.Context.NONE);
    }
}
