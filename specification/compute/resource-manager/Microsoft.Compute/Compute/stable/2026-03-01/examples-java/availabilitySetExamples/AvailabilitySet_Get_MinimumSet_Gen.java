
/**
 * Samples for AvailabilitySets GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/availabilitySetExamples/AvailabilitySet_Get_MinimumSet_Gen.json
     */
    /**
     * Sample code: AvailabilitySet_Get_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void availabilitySetGetMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getAvailabilitySets().getByResourceGroupWithResponse("rgcompute",
            "aaaaaaaaaaaaaaaaaaaa", com.azure.core.util.Context.NONE);
    }
}
