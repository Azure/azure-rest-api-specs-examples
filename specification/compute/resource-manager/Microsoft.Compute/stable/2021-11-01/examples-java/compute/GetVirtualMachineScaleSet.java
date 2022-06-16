import com.azure.core.util.Context;

/** Samples for VirtualMachineScaleSets GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/GetVirtualMachineScaleSet.json
     */
    /**
     * Sample code: Get a virtual machine scale set.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAVirtualMachineScaleSet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSets()
            .getByResourceGroupWithResponse("myResourceGroup", "myVirtualMachineScaleSet", null, Context.NONE);
    }
}
