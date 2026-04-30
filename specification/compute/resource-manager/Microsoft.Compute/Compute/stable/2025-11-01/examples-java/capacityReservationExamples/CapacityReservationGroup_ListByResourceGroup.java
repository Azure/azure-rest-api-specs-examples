
import com.azure.resourcemanager.compute.models.ExpandTypesForGetCapacityReservationGroups;

/**
 * Samples for CapacityReservationGroups ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/capacityReservationExamples/CapacityReservationGroup_ListByResourceGroup.json
     */
    /**
     * Sample code: List capacity reservation groups in resource group.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        listCapacityReservationGroupsInResourceGroup(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getCapacityReservationGroups().listByResourceGroup("myResourceGroup",
            ExpandTypesForGetCapacityReservationGroups.VIRTUAL_MACHINES_REF, com.azure.core.util.Context.NONE);
    }
}
