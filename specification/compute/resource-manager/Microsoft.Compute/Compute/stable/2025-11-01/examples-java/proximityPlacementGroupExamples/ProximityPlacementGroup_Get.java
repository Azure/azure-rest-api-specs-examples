
/**
 * Samples for ProximityPlacementGroups GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/proximityPlacementGroupExamples/ProximityPlacementGroup_Get.json
     */
    /**
     * Sample code: Get proximity placement groups.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getProximityPlacementGroups(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getProximityPlacementGroups().getByResourceGroupWithResponse("myResourceGroup",
            "myProximityPlacementGroup", null, com.azure.core.util.Context.NONE);
    }
}
