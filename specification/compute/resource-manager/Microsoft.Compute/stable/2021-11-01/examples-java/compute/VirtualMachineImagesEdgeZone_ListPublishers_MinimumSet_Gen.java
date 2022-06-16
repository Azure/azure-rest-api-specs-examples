import com.azure.core.util.Context;

/** Samples for VirtualMachineImagesEdgeZone ListPublishers. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineImagesEdgeZone_ListPublishers_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineImagesEdgeZone_ListPublishers_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineImagesEdgeZoneListPublishersMinimumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineImagesEdgeZones()
            .listPublishersWithResponse("aaaa", "aaaaaaaaaaaaaaaaaaaaaaa", Context.NONE);
    }
}
