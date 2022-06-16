import com.azure.core.util.Context;

/** Samples for CapacityReservations Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/capacityReservationExamples/CapacityReservation_Get.json
     */
    /**
     * Sample code: Get a capacity reservation.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getACapacityReservation(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getCapacityReservations()
            .getWithResponse(
                "myResourceGroup", "myCapacityReservationGroup", "myCapacityReservation", null, Context.NONE);
    }
}
