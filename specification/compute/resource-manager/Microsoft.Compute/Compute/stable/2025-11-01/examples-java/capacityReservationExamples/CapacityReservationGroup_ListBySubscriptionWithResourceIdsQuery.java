
import com.azure.resourcemanager.compute.models.ResourceIdOptionsForGetCapacityReservationGroups;

/**
 * Samples for CapacityReservationGroups List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2025-11-01/capacityReservationExamples/CapacityReservationGroup_ListBySubscriptionWithResourceIdsQuery.json
     */
    /**
     * Sample code: List capacity reservation groups with resource Ids only in subscription.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void listCapacityReservationGroupsWithResourceIdsOnlyInSubscription(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getCapacityReservationGroups().list(null,
            ResourceIdOptionsForGetCapacityReservationGroups.ALL, com.azure.core.util.Context.NONE);
    }
}
