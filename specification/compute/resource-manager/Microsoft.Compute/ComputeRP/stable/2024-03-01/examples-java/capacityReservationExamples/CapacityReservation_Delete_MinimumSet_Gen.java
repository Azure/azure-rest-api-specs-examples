
/**
 * Samples for CapacityReservations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/
     * capacityReservationExamples/CapacityReservation_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: CapacityReservation_Delete_MinimumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void capacityReservationDeleteMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getCapacityReservations().delete("rgcompute", "aaa", "aaaaaa",
            com.azure.core.util.Context.NONE);
    }
}
