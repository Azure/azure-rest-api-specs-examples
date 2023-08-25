/** Samples for ProximityPlacementGroups Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/proximityPlacementGroupExamples/ProximityPlacementGroup_Delete.json
     */
    /**
     * Sample code: Delete a proximity placement group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAProximityPlacementGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getProximityPlacementGroups()
            .deleteWithResponse("myResourceGroup", "myProximityPlacementGroup", com.azure.core.util.Context.NONE);
    }
}
