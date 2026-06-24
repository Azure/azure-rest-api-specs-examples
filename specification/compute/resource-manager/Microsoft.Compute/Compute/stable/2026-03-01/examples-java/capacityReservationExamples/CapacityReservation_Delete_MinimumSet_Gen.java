
/**
 * Samples for CapacityReservations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/capacityReservationExamples/CapacityReservation_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: CapacityReservation_Delete_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        capacityReservationDeleteMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getCapacityReservations().delete("rgcompute", "aaa", "aaaaaa",
            com.azure.core.util.Context.NONE);
    }
}
