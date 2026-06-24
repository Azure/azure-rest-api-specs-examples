
import com.azure.resourcemanager.compute.models.ExpandTypesForGetCapacityReservationGroups;

/**
 * Samples for CapacityReservationGroups List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/capacityReservationExamples/CapacityReservationGroup_ListBySubscription.json
     */
    /**
     * Sample code: List capacity reservation groups in subscription.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        listCapacityReservationGroupsInSubscription(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getCapacityReservationGroups().list(
            ExpandTypesForGetCapacityReservationGroups.VIRTUAL_MACHINES_REF, null, com.azure.core.util.Context.NONE);
    }
}
