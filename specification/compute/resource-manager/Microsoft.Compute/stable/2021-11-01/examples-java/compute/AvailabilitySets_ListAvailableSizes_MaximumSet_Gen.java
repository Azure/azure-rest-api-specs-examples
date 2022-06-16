import com.azure.core.util.Context;

/** Samples for AvailabilitySets ListAvailableSizes. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/AvailabilitySets_ListAvailableSizes_MaximumSet_Gen.json
     */
    /**
     * Sample code: AvailabilitySets_ListAvailableSizes_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void availabilitySetsListAvailableSizesMaximumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getAvailabilitySets()
            .listAvailableSizes("rgcompute", "aaaaaaaaaaaaaaaaaaaa", Context.NONE);
    }
}
