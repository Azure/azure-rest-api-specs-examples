
/**
 * Samples for CapacityReservationGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/capacityReservationExamples/CapacityReservationGroup_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: CapacityReservationGroup_Delete_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        capacityReservationGroupDeleteMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getCapacityReservationGroups().deleteWithResponse("rgcompute",
            "aaaaaaaaaaaaaaaaaaaaaaaaaa", com.azure.core.util.Context.NONE);
    }
}
