
/**
 * Samples for VirtualMachineExtensionImages Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/
     * virtualMachineExtensionImageExamples/VirtualMachineExtensionImage_Get_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineExtensionImage_Get_MinimumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        virtualMachineExtensionImageGetMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineExtensionImages().getWithResponse(
            "aaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaa", "aa", "aaa", com.azure.core.util.Context.NONE);
    }
}
