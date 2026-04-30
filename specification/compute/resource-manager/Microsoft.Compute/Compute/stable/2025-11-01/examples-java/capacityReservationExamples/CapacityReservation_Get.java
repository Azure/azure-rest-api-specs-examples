
import com.azure.resourcemanager.compute.models.CapacityReservationInstanceViewTypes;

/**
 * Samples for CapacityReservations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/capacityReservationExamples/CapacityReservation_Get.json
     */
    /**
     * Sample code: Get a capacity reservation.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getACapacityReservation(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getCapacityReservations().getWithResponse("myResourceGroup",
            "myCapacityReservationGroup", "myCapacityReservation", CapacityReservationInstanceViewTypes.INSTANCE_VIEW,
            com.azure.core.util.Context.NONE);
    }
}
