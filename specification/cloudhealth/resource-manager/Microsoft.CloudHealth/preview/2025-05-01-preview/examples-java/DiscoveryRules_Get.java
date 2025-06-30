
/**
 * Samples for DiscoveryRules Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01-preview/DiscoveryRules_Get.json
     */
    /**
     * Sample code: DiscoveryRules_Get.
     * 
     * @param manager Entry point to CloudHealthManager.
     */
    public static void discoveryRulesGet(com.azure.resourcemanager.cloudhealth.CloudHealthManager manager) {
        manager.discoveryRules().getWithResponse("myResourceGroup", "myHealthModel", "myDiscoveryRule",
            com.azure.core.util.Context.NONE);
    }
}
