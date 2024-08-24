
/**
 * Samples for VirtualMachineExtensionImages Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/
     * virtualMachineExtensionImageExamples/VirtualMachineExtensionImage_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineExtensionImage_Get_MaximumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        virtualMachineExtensionImageGetMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineExtensionImages().getWithResponse(
            "aaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaa",
            com.azure.core.util.Context.NONE);
    }
}
