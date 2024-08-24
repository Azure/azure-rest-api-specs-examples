
/**
 * Samples for VirtualMachineExtensionImages ListVersions.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/
     * virtualMachineExtensionImageExamples/VirtualMachineExtensionImage_ListVersions_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineExtensionImage_ListVersions_MinimumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        virtualMachineExtensionImageListVersionsMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineExtensionImages().listVersionsWithResponse(
            "aaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaa", null, null, null, com.azure.core.util.Context.NONE);
    }
}
