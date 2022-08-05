import com.azure.core.util.Context;

/** Samples for VirtualMachineScaleSetVMExtensions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/DeleteVirtualMachineScaleSetVMExtensions.json
     */
    /**
     * Sample code: Delete VirtualMachineScaleSet VM extension.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteVirtualMachineScaleSetVMExtension(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSetVMExtensions()
            .delete("myResourceGroup", "myvmScaleSet", "0", "myVMExtension", Context.NONE);
    }
}
