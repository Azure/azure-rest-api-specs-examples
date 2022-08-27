import com.azure.core.util.Context;

/** Samples for VirtualMachines ListAvailableSizes. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/virtualMachineExamples/VirtualMachines_ListAvailableVmSizes.json
     */
    /**
     * Sample code: Lists all available virtual machine sizes to which the specified virtual machine can be resized.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listsAllAvailableVirtualMachineSizesToWhichTheSpecifiedVirtualMachineCanBeResized(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachines()
            .listAvailableSizes("myResourceGroup", "myVmName", Context.NONE);
    }
}
