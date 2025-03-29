
/**
 * Samples for AvailabilitySets ListAvailableSizes.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/
     * availabilitySetExamples/AvailabilitySet_ListAvailableSizes_MaximumSet_Gen.json
     */
    /**
     * Sample code: AvailabilitySet_ListAvailableSizes_MaximumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        availabilitySetListAvailableSizesMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getAvailabilitySets().listAvailableSizes("rgcompute",
            "aaaaaaaaaaaaaaaaaaaa", com.azure.core.util.Context.NONE);
    }
}
