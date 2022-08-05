import com.azure.core.util.Context;

/** Samples for VirtualMachineScaleSets Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/compute/ForceDeleteVirtualMachineScaleSets.json
     */
    /**
     * Sample code: Force Delete a VM scale set.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void forceDeleteAVMScaleSet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSets()
            .delete("myResourceGroup", "myvmScaleSet", true, Context.NONE);
    }
}
