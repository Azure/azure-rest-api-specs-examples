
/**
 * Samples for VirtualMachineScaleSetVMExtensions Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/
     * virtualMachineScaleSetExamples/VirtualMachineScaleSetVMExtension_Get.json
     */
    /**
     * Sample code: Get VirtualMachineScaleSet VM extension.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getVirtualMachineScaleSetVMExtension(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSetVMExtensions().getWithResponse(
            "myResourceGroup", "myvmScaleSet", "0", "myVMExtension", null, com.azure.core.util.Context.NONE);
    }
}
