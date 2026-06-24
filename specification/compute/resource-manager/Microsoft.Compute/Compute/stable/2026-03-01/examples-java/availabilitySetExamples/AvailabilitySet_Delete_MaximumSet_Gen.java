
/**
 * Samples for AvailabilitySets Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/availabilitySetExamples/AvailabilitySet_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: AvailabilitySet_Delete_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void availabilitySetDeleteMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getAvailabilitySets().deleteWithResponse("rgcompute", "aaaaaaaaaaaaaaaaaaaa",
            com.azure.core.util.Context.NONE);
    }
}
