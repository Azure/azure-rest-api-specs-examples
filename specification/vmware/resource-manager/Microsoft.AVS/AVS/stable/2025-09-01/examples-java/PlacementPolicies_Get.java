
/**
 * Samples for PlacementPolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/PlacementPolicies_Get.json
     */
    /**
     * Sample code: PlacementPolicies_Get.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void placementPoliciesGet(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.placementPolicies().getWithResponse("group1", "cloud1", "cluster1", "policy1",
            com.azure.core.util.Context.NONE);
    }
}
