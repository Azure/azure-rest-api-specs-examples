
import com.azure.resourcemanager.compute.models.CapacityReservationInstanceViewTypes;

/**
 * Samples for CapacityReservations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/capacityReservationExamples/BlockCapacityReservation_Get.json
     */
    /**
     * Sample code: Get a block capacity reservation.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getABlockCapacityReservation(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getCapacityReservations().getWithResponse("myResourceGroup",
            "blockCapacityReservationGroup", "blockCapacityReservation",
            CapacityReservationInstanceViewTypes.INSTANCE_VIEW, com.azure.core.util.Context.NONE);
    }
}
