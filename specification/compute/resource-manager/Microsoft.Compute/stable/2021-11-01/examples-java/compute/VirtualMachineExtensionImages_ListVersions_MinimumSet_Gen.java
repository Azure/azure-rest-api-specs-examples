import com.azure.core.util.Context;

/** Samples for VirtualMachineExtensionImages ListVersions. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineExtensionImages_ListVersions_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineExtensionImages_ListVersions_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineExtensionImagesListVersionsMinimumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineExtensionImages()
            .listVersionsWithResponse(
                "aaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaa", null, null, null, Context.NONE);
    }
}
