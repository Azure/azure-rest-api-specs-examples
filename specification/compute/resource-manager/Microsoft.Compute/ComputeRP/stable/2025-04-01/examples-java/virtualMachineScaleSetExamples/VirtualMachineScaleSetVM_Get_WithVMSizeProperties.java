
/**
 * Samples for VirtualMachineScaleSetVMs Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/
     * virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_Get_WithVMSizeProperties.json
     */
    /**
     * Sample code: Get VM scale set VM with VMSizeProperties.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getVMScaleSetVMWithVMSizeProperties(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSetVMs()
            .getWithResponse("myResourceGroup", "{vmss-name}", "0", null, com.azure.core.util.Context.NONE);
    }
}
