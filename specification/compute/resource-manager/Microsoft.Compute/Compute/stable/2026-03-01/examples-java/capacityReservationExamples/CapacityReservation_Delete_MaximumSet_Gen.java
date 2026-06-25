
/**
 * Samples for CapacityReservations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/capacityReservationExamples/CapacityReservation_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: CapacityReservation_Delete_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        capacityReservationDeleteMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getCapacityReservations().delete("rgcompute", "aaaaaaaaaaa",
            "aaaaaaaaaaaaaaaaaaaaaaaaaaa", com.azure.core.util.Context.NONE);
    }
}
