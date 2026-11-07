
/**
 * Samples for NeuroShieldSummaries List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/NeuroShieldSummaries_List.json
     */
    /**
     * Sample code: List NeuroShield fleet summaries.
     * 
     * @param manager Entry point to NetworkSecurityManager.
     */
    public static void
        listNeuroShieldFleetSummaries(com.azure.resourcemanager.networksecurity.NetworkSecurityManager manager) {
        manager.neuroShieldSummaries().list("myResourceGroup", "myNeuroShield", null, com.azure.core.util.Context.NONE);
    }
}
