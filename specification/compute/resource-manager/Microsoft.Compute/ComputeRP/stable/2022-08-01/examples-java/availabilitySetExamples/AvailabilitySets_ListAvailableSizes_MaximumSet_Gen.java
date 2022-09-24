import com.azure.core.util.Context;

/** Samples for AvailabilitySets ListAvailableSizes. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/availabilitySetExamples/AvailabilitySets_ListAvailableSizes_MaximumSet_Gen.json
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
