
/**
 * Samples for NeuroShieldSummaries Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/NeuroShieldSummaries_Get.json
     */
    /**
     * Sample code: Get a NeuroShield fleet summary.
     * 
     * @param manager Entry point to NetworkSecurityManager.
     */
    public static void
        getANeuroShieldFleetSummary(com.azure.resourcemanager.networksecurity.NetworkSecurityManager manager) {
        manager.neuroShieldSummaries().getWithResponse("myResourceGroup", "myNeuroShield",
            com.azure.core.util.Context.NONE);
    }
}
