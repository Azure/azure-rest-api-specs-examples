
/**
 * Samples for DiscoveryRules ListByHealthModel.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01-preview/DiscoveryRules_ListByHealthModel.json
     */
    /**
     * Sample code: DiscoveryRules_ListByHealthModel.
     * 
     * @param manager Entry point to CloudHealthManager.
     */
    public static void
        discoveryRulesListByHealthModel(com.azure.resourcemanager.cloudhealth.CloudHealthManager manager) {
        manager.discoveryRules().listByHealthModel("my-resource-group", "my-health-model", null,
            com.azure.core.util.Context.NONE);
    }
}
