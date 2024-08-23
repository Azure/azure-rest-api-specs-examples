
import com.azure.resourcemanager.compute.models.ExpandTypesForGetCapacityReservationGroups;

/**
 * Samples for CapacityReservationGroups List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/
     * capacityReservationExamples/CapacityReservationGroup_ListBySubscription.json
     */
    /**
     * Sample code: List capacity reservation groups in subscription.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listCapacityReservationGroupsInSubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getCapacityReservationGroups().list(
            ExpandTypesForGetCapacityReservationGroups.VIRTUAL_MACHINES_REF, null, com.azure.core.util.Context.NONE);
    }
}
