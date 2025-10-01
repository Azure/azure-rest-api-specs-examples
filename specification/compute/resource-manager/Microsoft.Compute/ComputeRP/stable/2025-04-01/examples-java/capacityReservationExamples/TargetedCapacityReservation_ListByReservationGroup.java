
/**
 * Samples for CapacityReservations ListByCapacityReservationGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/
     * capacityReservationExamples/TargetedCapacityReservation_ListByReservationGroup.json
     */
    /**
     * Sample code: List capacity reservations in targeted reservation group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listCapacityReservationsInTargetedReservationGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getCapacityReservations().listByCapacityReservationGroup(
            "myResourceGroup", "targetedCapacityReservationGroup", com.azure.core.util.Context.NONE);
    }
}
