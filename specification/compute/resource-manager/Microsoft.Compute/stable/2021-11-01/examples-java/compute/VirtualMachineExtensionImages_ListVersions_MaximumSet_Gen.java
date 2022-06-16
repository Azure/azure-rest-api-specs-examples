import com.azure.core.util.Context;

/** Samples for VirtualMachineExtensionImages ListVersions. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineExtensionImages_ListVersions_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineExtensionImages_ListVersions_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineExtensionImagesListVersionsMaximumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineExtensionImages()
            .listVersionsWithResponse(
                "aaaaaaaaaaaaaaaaaaaaaaaaaa",
                "aaaaaaaaaaaaaaaaaaaa",
                "aaaaaaaaaaaaaaaaaa",
                "aaaaaaaaaaaaaaaaaaaaaaaaa",
                22,
                "a",
                Context.NONE);
    }
}
