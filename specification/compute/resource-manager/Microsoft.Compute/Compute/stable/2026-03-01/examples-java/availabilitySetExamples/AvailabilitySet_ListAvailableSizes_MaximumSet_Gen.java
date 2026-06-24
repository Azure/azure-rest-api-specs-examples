
/**
 * Samples for AvailabilitySets ListAvailableSizes.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/availabilitySetExamples/AvailabilitySet_ListAvailableSizes_MaximumSet_Gen.json
     */
    /**
     * Sample code: AvailabilitySet_ListAvailableSizes_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        availabilitySetListAvailableSizesMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getAvailabilitySets().listAvailableSizes("rgcompute", "aaaaaaaaaaaaaaaaaaaa",
            com.azure.core.util.Context.NONE);
    }
}
