import com.azure.core.util.Context;

/** Samples for VirtualMachineImages ListPublishers. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineImages_ListPublishers_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineImages_ListPublishers_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineImagesListPublishersMinimumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineImages()
            .listPublishersWithResponse("aaaaaaaaaaaaaaaaaaaaaaaaaaa", Context.NONE);
    }
}
