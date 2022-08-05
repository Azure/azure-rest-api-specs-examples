import com.azure.core.util.Context;

/** Samples for CapacityReservations ListByCapacityReservationGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/ListCapacityReservationsInReservationGroup.json
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
            .listByCapacityReservationGroup("myResourceGroup", "myCapacityReservationGroup", Context.NONE);
    }
}
