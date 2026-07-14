
/**
 * Samples for DiscoveryRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/DiscoveryRules_Delete.json
     */
    /**
     * Sample code: DiscoveryRules_Delete.
     * 
     * @param manager Entry point to CloudHealthManager.
     */
    public static void discoveryRulesDelete(com.azure.resourcemanager.cloudhealth.CloudHealthManager manager) {
        manager.discoveryRules().delete("online-store-rg", "online-store", "discover-web-apps",
            com.azure.core.util.Context.NONE);
    }
}
