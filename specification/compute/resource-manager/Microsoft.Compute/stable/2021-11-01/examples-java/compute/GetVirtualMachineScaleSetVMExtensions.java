import com.azure.core.util.Context;

/** Samples for VirtualMachineScaleSetVMExtensions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/GetVirtualMachineScaleSetVMExtensions.json
     */
    /**
     * Sample code: Get VirtualMachineScaleSet VM extension.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getVirtualMachineScaleSetVMExtension(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSetVMExtensions()
            .getWithResponse("myResourceGroup", "myvmScaleSet", "0", "myVMExtension", null, Context.NONE);
    }
}
