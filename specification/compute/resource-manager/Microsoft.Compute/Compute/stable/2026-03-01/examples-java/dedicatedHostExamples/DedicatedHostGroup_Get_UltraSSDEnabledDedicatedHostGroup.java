
/**
 * Samples for DedicatedHostGroups GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2026-03-01/dedicatedHostExamples/DedicatedHostGroup_Get_UltraSSDEnabledDedicatedHostGroup.json
     */
    /**
     * Sample code: Create an ultraSSDEnabled dedicated host group.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        createAnUltraSSDEnabledDedicatedHostGroup(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDedicatedHostGroups().getByResourceGroupWithResponse("myResourceGroup",
            "myDedicatedHostGroup", null, com.azure.core.util.Context.NONE);
    }
}
