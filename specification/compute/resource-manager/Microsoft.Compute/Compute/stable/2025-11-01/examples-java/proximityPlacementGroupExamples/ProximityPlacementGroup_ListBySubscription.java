
/**
 * Samples for ProximityPlacementGroups List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/proximityPlacementGroupExamples/ProximityPlacementGroup_ListBySubscription.json
     */
    /**
     * Sample code: List proximity placement groups.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void listProximityPlacementGroups(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getProximityPlacementGroups().list(com.azure.core.util.Context.NONE);
    }
}
