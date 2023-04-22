/** Samples for VirtualMachineExtensionImages ListTypes. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-03-01/examples/virtualMachineExtensionImageExamples/VirtualMachineExtensionImage_ListTypes_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineExtensionImage_ListTypes_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineExtensionImageListTypesMinimumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineExtensionImages()
            .listTypesWithResponse("aaaa", "aa", com.azure.core.util.Context.NONE);
    }
}
