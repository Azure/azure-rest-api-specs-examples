
import com.azure.resourcemanager.compute.models.CapacityReservationInstanceViewTypes;

/**
 * Samples for CapacityReservations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/capacityReservationExamples/OpenCapacityReservation_Get.json
     */
    /**
     * Sample code: Get an open capacity reservation with instance view.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        getAnOpenCapacityReservationWithInstanceView(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getCapacityReservations().getWithResponse("myResourceGroup",
            "openCapacityReservationGroup", "openCapacityReservation",
            CapacityReservationInstanceViewTypes.INSTANCE_VIEW, com.azure.core.util.Context.NONE);
    }
}
