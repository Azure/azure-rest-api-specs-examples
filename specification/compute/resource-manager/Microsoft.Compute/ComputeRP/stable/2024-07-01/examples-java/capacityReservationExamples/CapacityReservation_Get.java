
import com.azure.resourcemanager.compute.models.CapacityReservationInstanceViewTypes;

/**
 * Samples for CapacityReservations Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/
     * capacityReservationExamples/CapacityReservation_Get.json
     */
    /**
     * Sample code: Get a capacity reservation.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getACapacityReservation(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getCapacityReservations().getWithResponse("myResourceGroup",
            "myCapacityReservationGroup", "myCapacityReservation", CapacityReservationInstanceViewTypes.INSTANCE_VIEW,
            com.azure.core.util.Context.NONE);
    }
}
