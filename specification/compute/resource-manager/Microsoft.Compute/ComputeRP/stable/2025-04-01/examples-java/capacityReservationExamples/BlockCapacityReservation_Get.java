
import com.azure.resourcemanager.compute.models.CapacityReservationInstanceViewTypes;

/**
 * Samples for CapacityReservations Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/
     * capacityReservationExamples/BlockCapacityReservation_Get.json
     */
    /**
     * Sample code: Get a block capacity reservation.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getABlockCapacityReservation(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getCapacityReservations().getWithResponse("myResourceGroup",
            "blockCapacityReservationGroup", "blockCapacityReservation",
            CapacityReservationInstanceViewTypes.INSTANCE_VIEW, com.azure.core.util.Context.NONE);
    }
}
