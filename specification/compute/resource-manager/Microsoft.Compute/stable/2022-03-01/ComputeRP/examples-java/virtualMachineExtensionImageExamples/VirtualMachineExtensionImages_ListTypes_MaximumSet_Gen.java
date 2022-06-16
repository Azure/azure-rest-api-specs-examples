import com.azure.core.util.Context;

/** Samples for VirtualMachineExtensionImages ListTypes. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineExtensionImageExamples/VirtualMachineExtensionImages_ListTypes_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineExtensionImages_ListTypes_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineExtensionImagesListTypesMaximumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineExtensionImages()
            .listTypesWithResponse("aaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", Context.NONE);
    }
}
