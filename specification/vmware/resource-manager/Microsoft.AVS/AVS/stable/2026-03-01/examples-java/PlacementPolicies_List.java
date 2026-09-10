
/**
 * Samples for PlacementPolicies List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/PlacementPolicies_List.json
     */
    /**
     * Sample code: PlacementPolicies_List.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void placementPoliciesList(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.placementPolicies().list("group1", "cloud1", "cluster1", com.azure.core.util.Context.NONE);
    }
}
