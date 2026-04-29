
/**
 * Samples for AvailabilitySets Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/availabilitySetExamples/AvailabilitySet_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: AvailabilitySet_Delete_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void availabilitySetDeleteMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getAvailabilitySets().deleteWithResponse("rgcompute", "aaaaaaaaaaa",
            com.azure.core.util.Context.NONE);
    }
}
