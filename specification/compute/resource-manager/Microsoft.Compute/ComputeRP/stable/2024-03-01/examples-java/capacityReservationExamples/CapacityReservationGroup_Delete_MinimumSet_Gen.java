
/**
 * Samples for CapacityReservationGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/
     * capacityReservationExamples/CapacityReservationGroup_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: CapacityReservationGroup_Delete_MinimumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        capacityReservationGroupDeleteMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getCapacityReservationGroups().deleteWithResponse("rgcompute",
            "aaaaaaaaaaaaaaaaaaaaaaaaaa", com.azure.core.util.Context.NONE);
    }
}
