
/**
 * Samples for DiscoveryRules Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/DiscoveryRules_Get.json
     */
    /**
     * Sample code: DiscoveryRules_Get.
     * 
     * @param manager Entry point to CloudHealthManager.
     */
    public static void discoveryRulesGet(com.azure.resourcemanager.cloudhealth.CloudHealthManager manager) {
        manager.discoveryRules().getWithResponse("online-store-rg", "online-store", "discover-web-apps",
            com.azure.core.util.Context.NONE);
    }
}
