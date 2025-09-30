
/**
 * Samples for CapacityReservations ListByCapacityReservationGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/
     * capacityReservationExamples/BlockCapacityReservation_ListByReservationGroup.json
     */
    /**
     * Sample code: List block capacity reservations in reservation group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listBlockCapacityReservationsInReservationGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getCapacityReservations().listByCapacityReservationGroup(
            "myResourceGroup", "blockCapacityReservationGroup", com.azure.core.util.Context.NONE);
    }
}
