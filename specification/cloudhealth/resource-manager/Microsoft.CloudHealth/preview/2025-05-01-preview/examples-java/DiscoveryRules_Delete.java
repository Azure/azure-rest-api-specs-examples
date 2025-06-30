
/**
 * Samples for DiscoveryRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01-preview/DiscoveryRules_Delete.json
     */
    /**
     * Sample code: DiscoveryRules_Delete.
     * 
     * @param manager Entry point to CloudHealthManager.
     */
    public static void discoveryRulesDelete(com.azure.resourcemanager.cloudhealth.CloudHealthManager manager) {
        manager.discoveryRules().deleteWithResponse("my-resource-group", "my-health-model", "my-discovery-rule",
            com.azure.core.util.Context.NONE);
    }
}
