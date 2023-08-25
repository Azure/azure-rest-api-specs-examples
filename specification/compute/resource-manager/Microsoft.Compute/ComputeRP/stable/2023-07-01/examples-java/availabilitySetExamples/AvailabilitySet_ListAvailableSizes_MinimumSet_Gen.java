/** Samples for AvailabilitySets ListAvailableSizes. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/availabilitySetExamples/AvailabilitySet_ListAvailableSizes_MinimumSet_Gen.json
     */
    /**
     * Sample code: AvailabilitySet_ListAvailableSizes_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void availabilitySetListAvailableSizesMinimumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getAvailabilitySets()
            .listAvailableSizes("rgcompute", "aa", com.azure.core.util.Context.NONE);
    }
}
