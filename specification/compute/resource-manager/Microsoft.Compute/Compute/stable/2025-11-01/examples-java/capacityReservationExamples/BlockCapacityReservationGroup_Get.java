
import com.azure.resourcemanager.compute.models.CapacityReservationGroupInstanceViewTypes;

/**
 * Samples for CapacityReservationGroups GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/capacityReservationExamples/BlockCapacityReservationGroup_Get.json
     */
    /**
     * Sample code: Get a block capacity reservation Group.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getABlockCapacityReservationGroup(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getCapacityReservationGroups().getByResourceGroupWithResponse("myResourceGroup",
            "blockCapacityReservationGroup", CapacityReservationGroupInstanceViewTypes.INSTANCE_VIEW,
            com.azure.core.util.Context.NONE);
    }
}
