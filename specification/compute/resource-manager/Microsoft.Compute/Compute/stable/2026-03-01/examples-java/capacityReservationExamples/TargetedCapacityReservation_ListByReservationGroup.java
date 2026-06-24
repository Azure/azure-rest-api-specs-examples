
import com.azure.resourcemanager.compute.models.ExpandTypesForGetCapacityReservationGroups;

/**
 * Samples for CapacityReservations ListByCapacityReservationGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2026-03-01/capacityReservationExamples/TargetedCapacityReservation_ListByReservationGroup.json
     */
    /**
     * Sample code: List capacity reservations in targeted reservation group.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        listCapacityReservationsInTargetedReservationGroup(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getCapacityReservations().listByCapacityReservationGroup("myResourceGroup",
            "targetedCapacityReservationGroup", ExpandTypesForGetCapacityReservationGroups.VIRTUAL_MACHINES_REF,
            com.azure.core.util.Context.NONE);
    }
}
