import com.azure.core.util.Context;

/** Samples for VirtualMachineImagesEdgeZone ListSkus. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineImagesEdgeZone_ListSkus_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineImagesEdgeZone_ListSkus_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineImagesEdgeZoneListSkusMinimumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineImagesEdgeZones()
            .listSkusWithResponse(
                "aaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaa", "aaaaaaaaaaaa", Context.NONE);
    }
}
