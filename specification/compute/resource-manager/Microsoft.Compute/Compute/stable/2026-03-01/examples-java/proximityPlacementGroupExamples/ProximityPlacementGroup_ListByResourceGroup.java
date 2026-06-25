
/**
 * Samples for ProximityPlacementGroups ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/proximityPlacementGroupExamples/ProximityPlacementGroup_ListByResourceGroup.json
     */
    /**
     * Sample code: List proximity placement group.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void listProximityPlacementGroup(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getProximityPlacementGroups().listByResourceGroup("myResourceGroup",
            com.azure.core.util.Context.NONE);
    }
}
