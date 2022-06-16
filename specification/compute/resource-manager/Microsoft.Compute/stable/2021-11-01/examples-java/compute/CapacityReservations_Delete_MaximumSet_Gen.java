import com.azure.core.util.Context;

/** Samples for CapacityReservations Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/CapacityReservations_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: CapacityReservations_Delete_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void capacityReservationsDeleteMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getCapacityReservations()
            .delete("rgcompute", "aaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaa", Context.NONE);
    }
}
