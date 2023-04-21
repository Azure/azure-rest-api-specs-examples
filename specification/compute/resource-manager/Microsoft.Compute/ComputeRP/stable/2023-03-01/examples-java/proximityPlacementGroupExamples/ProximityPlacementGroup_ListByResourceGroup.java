/** Samples for ProximityPlacementGroups ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-03-01/examples/proximityPlacementGroupExamples/ProximityPlacementGroup_ListByResourceGroup.json
     */
    /**
     * Sample code: List proximity placement group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listProximityPlacementGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getProximityPlacementGroups()
            .listByResourceGroup("myResourceGroup", com.azure.core.util.Context.NONE);
    }
}
