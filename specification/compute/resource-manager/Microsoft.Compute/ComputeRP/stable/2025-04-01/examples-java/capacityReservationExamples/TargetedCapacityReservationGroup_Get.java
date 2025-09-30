
import com.azure.resourcemanager.compute.models.CapacityReservationGroupInstanceViewTypes;

/**
 * Samples for CapacityReservationGroups GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/
     * capacityReservationExamples/TargetedCapacityReservationGroup_Get.json
     */
    /**
     * Sample code: Get a targeted capacity reservation group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getATargetedCapacityReservationGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getCapacityReservationGroups().getByResourceGroupWithResponse(
            "myResourceGroup", "targetedCapacityReservationGroup",
            CapacityReservationGroupInstanceViewTypes.INSTANCE_VIEW, com.azure.core.util.Context.NONE);
    }
}
