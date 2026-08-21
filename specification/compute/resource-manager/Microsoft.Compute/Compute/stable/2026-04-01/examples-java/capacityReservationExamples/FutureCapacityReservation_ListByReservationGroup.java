
import com.azure.resourcemanager.compute.models.ExpandTypesForGetCapacityReservationGroups;

/**
 * Samples for CapacityReservations ListByCapacityReservationGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/capacityReservationExamples/FutureCapacityReservation_ListByReservationGroup.json
     */
    /**
     * Sample code: List Future capacity reservations in reservation group.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        listFutureCapacityReservationsInReservationGroup(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getCapacityReservations().listByCapacityReservationGroup("myResourceGroup",
            "futureCapacityReservationGroup", ExpandTypesForGetCapacityReservationGroups.VIRTUAL_MACHINES_REF,
            com.azure.core.util.Context.NONE);
    }
}
