
import com.azure.resourcemanager.compute.models.CapacityReservationGroupInstanceViewTypes;

/**
 * Samples for CapacityReservationGroups GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/capacityReservationExamples/TargetedCapacityReservationGroup_Get.json
     */
    /**
     * Sample code: Get a targeted capacity reservation group.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getATargetedCapacityReservationGroup(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getCapacityReservationGroups().getByResourceGroupWithResponse("myResourceGroup",
            "targetedCapacityReservationGroup", CapacityReservationGroupInstanceViewTypes.INSTANCE_VIEW,
            com.azure.core.util.Context.NONE);
    }
}
