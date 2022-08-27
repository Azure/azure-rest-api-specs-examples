import com.azure.core.util.Context;

/** Samples for VirtualMachineExtensionImages ListTypes. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/virtualMachineExtensionImageExamples/VirtualMachineExtensionImages_ListTypes_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineExtensionImages_ListTypes_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineExtensionImagesListTypesMinimumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineExtensionImages()
            .listTypesWithResponse("aaaa", "aa", Context.NONE);
    }
}
