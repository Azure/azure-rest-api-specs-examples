
import com.azure.resourcemanager.compute.models.CapacityReservationInstanceViewTypes;

/**
 * Samples for CapacityReservations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/capacityReservationExamples/TargetedCapacityReservation_Get.json
     */
    /**
     * Sample code: Get a targeted capacity reservation.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getATargetedCapacityReservation(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getCapacityReservations().getWithResponse("myResourceGroup",
            "targetedCapacityReservationGroup", "targetedCapacityReservation",
            CapacityReservationInstanceViewTypes.INSTANCE_VIEW, com.azure.core.util.Context.NONE);
    }
}
