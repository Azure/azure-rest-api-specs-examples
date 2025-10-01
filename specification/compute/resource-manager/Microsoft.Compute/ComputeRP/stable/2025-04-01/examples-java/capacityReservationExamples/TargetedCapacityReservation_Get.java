
import com.azure.resourcemanager.compute.models.CapacityReservationInstanceViewTypes;

/**
 * Samples for CapacityReservations Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/
     * capacityReservationExamples/TargetedCapacityReservation_Get.json
     */
    /**
     * Sample code: Get a targeted capacity reservation.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getATargetedCapacityReservation(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getCapacityReservations().getWithResponse("myResourceGroup",
            "targetedCapacityReservationGroup", "targetedCapacityReservation",
            CapacityReservationInstanceViewTypes.INSTANCE_VIEW, com.azure.core.util.Context.NONE);
    }
}
