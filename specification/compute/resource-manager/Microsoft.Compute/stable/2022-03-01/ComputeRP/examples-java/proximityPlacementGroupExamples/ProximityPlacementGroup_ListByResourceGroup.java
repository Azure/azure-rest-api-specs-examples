import com.azure.core.util.Context;

/** Samples for ProximityPlacementGroups ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/proximityPlacementGroupExamples/ProximityPlacementGroup_ListByResourceGroup.json
     */
    /**
     * Sample code: Create a proximity placement group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAProximityPlacementGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getProximityPlacementGroups()
            .listByResourceGroup("myResourceGroup", Context.NONE);
    }
}
