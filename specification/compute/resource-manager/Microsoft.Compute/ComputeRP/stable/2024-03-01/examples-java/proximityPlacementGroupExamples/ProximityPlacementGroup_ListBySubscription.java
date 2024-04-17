
/**
 * Samples for ProximityPlacementGroups List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/
     * proximityPlacementGroupExamples/ProximityPlacementGroup_ListBySubscription.json
     */
    /**
     * Sample code: List proximity placement groups.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listProximityPlacementGroups(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getProximityPlacementGroups()
            .list(com.azure.core.util.Context.NONE);
    }
}
