import com.azure.core.util.Context;

/** Samples for CapacityReservationGroups GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/GetACapacityReservationGroup.json
     */
    /**
     * Sample code: Get a capacity reservation Group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getACapacityReservationGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getCapacityReservationGroups()
            .getByResourceGroupWithResponse("myResourceGroup", "myCapacityReservationGroup", null, Context.NONE);
    }
}
