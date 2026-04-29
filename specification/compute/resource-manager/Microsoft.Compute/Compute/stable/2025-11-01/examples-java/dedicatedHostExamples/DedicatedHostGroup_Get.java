
/**
 * Samples for DedicatedHostGroups GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/dedicatedHostExamples/DedicatedHostGroup_Get.json
     */
    /**
     * Sample code: Create a dedicated host group.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void createADedicatedHostGroup(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDedicatedHostGroups().getByResourceGroupWithResponse("myResourceGroup",
            "myDedicatedHostGroup", null, com.azure.core.util.Context.NONE);
    }
}
