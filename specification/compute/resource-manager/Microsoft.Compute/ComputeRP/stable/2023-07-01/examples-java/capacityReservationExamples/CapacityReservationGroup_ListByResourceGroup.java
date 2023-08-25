import com.azure.resourcemanager.compute.models.ExpandTypesForGetCapacityReservationGroups;

/** Samples for CapacityReservationGroups ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/capacityReservationExamples/CapacityReservationGroup_ListByResourceGroup.json
     */
    /**
     * Sample code: List capacity reservation groups in resource group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listCapacityReservationGroupsInResourceGroup(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getCapacityReservationGroups()
            .listByResourceGroup(
                "myResourceGroup",
                ExpandTypesForGetCapacityReservationGroups.VIRTUAL_MACHINES_REF,
                com.azure.core.util.Context.NONE);
    }
}
