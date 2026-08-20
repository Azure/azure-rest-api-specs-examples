
import com.azure.resourcemanager.compute.models.CapacityReservationInstanceViewTypes;

/**
 * Samples for CapacityReservations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/capacityReservationExamples/FutureCapacityReservation_Get.json
     */
    /**
     * Sample code: Get a Future capacity reservation.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getAFutureCapacityReservation(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getCapacityReservations().getWithResponse("myResourceGroup",
            "futureCapacityReservationGroup", "futureCapacityReservation",
            CapacityReservationInstanceViewTypes.INSTANCE_VIEW, com.azure.core.util.Context.NONE);
    }
}
