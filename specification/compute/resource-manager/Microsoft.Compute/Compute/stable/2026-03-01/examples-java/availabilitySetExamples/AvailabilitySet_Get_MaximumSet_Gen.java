
/**
 * Samples for AvailabilitySets GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/availabilitySetExamples/AvailabilitySet_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: AvailabilitySet_Get_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void availabilitySetGetMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getAvailabilitySets().getByResourceGroupWithResponse("rgcompute", "aaaaaaaaaaaa",
            com.azure.core.util.Context.NONE);
    }
}
