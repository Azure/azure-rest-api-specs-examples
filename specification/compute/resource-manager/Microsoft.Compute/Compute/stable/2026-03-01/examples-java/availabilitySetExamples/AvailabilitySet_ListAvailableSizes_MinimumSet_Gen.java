
/**
 * Samples for AvailabilitySets ListAvailableSizes.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/availabilitySetExamples/AvailabilitySet_ListAvailableSizes_MinimumSet_Gen.json
     */
    /**
     * Sample code: AvailabilitySet_ListAvailableSizes_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        availabilitySetListAvailableSizesMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getAvailabilitySets().listAvailableSizes("rgcompute", "aa",
            com.azure.core.util.Context.NONE);
    }
}
