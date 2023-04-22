/** Samples for CapacityReservations ListByCapacityReservationGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-03-01/examples/capacityReservationExamples/CapacityReservation_ListByReservationGroup.json
     */
    /**
     * Sample code: List capacity reservations in reservation group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listCapacityReservationsInReservationGroup(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getCapacityReservations()
            .listByCapacityReservationGroup(
                "myResourceGroup", "myCapacityReservationGroup", com.azure.core.util.Context.NONE);
    }
}
