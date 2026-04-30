
/**
 * Samples for ProximityPlacementGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/proximityPlacementGroupExamples/ProximityPlacementGroup_Delete.json
     */
    /**
     * Sample code: Delete a proximity placement group.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void deleteAProximityPlacementGroup(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getProximityPlacementGroups().deleteWithResponse("myResourceGroup",
            "myProximityPlacementGroup", com.azure.core.util.Context.NONE);
    }
}
