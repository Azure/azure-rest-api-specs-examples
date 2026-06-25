
/**
 * Samples for AvailabilitySets ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/availabilitySetExamples/AvailabilitySet_List_MinimumSet_Gen.json
     */
    /**
     * Sample code: AvailabilitySet_List_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void availabilitySetListMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getAvailabilitySets().listByResourceGroup("rgcompute",
            com.azure.core.util.Context.NONE);
    }
}
