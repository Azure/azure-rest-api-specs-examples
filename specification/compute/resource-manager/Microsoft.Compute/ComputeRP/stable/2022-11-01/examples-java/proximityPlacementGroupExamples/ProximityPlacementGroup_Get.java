/** Samples for ProximityPlacementGroups GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-11-01/examples/proximityPlacementGroupExamples/ProximityPlacementGroup_Get.json
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
            .getByResourceGroupWithResponse(
                "myResourceGroup", "myProximityPlacementGroup", null, com.azure.core.util.Context.NONE);
    }
}
