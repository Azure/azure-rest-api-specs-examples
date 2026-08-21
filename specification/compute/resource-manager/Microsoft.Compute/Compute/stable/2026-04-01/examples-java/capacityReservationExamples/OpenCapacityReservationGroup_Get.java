
import com.azure.resourcemanager.compute.models.CapacityReservationGroupInstanceViewTypes;

/**
 * Samples for CapacityReservationGroups GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/capacityReservationExamples/OpenCapacityReservationGroup_Get.json
     */
    /**
     * Sample code: Get an open capacity reservation group with instance view.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        getAnOpenCapacityReservationGroupWithInstanceView(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getCapacityReservationGroups().getByResourceGroupWithResponse("myResourceGroup",
            "openCapacityReservationGroup", CapacityReservationGroupInstanceViewTypes.INSTANCE_VIEW,
            com.azure.core.util.Context.NONE);
    }
}
