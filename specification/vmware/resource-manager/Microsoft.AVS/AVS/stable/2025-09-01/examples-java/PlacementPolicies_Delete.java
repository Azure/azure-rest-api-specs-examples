
/**
 * Samples for PlacementPolicies Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/PlacementPolicies_Delete.json
     */
    /**
     * Sample code: PlacementPolicies_Delete.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void placementPoliciesDelete(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.placementPolicies().delete("group1", "cloud1", "cluster1", "policy1", com.azure.core.util.Context.NONE);
    }
}
