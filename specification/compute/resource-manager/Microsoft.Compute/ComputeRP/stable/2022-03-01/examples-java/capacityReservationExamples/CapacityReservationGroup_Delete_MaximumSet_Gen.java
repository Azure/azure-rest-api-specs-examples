import com.azure.core.util.Context;

/** Samples for CapacityReservationGroups Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/capacityReservationExamples/CapacityReservationGroup_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: CapacityReservationGroups_Delete_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void capacityReservationGroupsDeleteMaximumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getCapacityReservationGroups()
            .deleteWithResponse("rgcompute", "a", Context.NONE);
    }
}
