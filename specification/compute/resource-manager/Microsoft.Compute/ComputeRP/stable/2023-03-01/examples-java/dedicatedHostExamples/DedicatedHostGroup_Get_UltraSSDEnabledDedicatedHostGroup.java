/** Samples for DedicatedHostGroups GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-03-01/examples/dedicatedHostExamples/DedicatedHostGroup_Get_UltraSSDEnabledDedicatedHostGroup.json
     */
    /**
     * Sample code: Create an ultraSSDEnabled dedicated host group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAnUltraSSDEnabledDedicatedHostGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getDedicatedHostGroups()
            .getByResourceGroupWithResponse(
                "myResourceGroup", "myDedicatedHostGroup", null, com.azure.core.util.Context.NONE);
    }
}
