
/**
 * Samples for DiscoveryRules ListByHealthModel.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/DiscoveryRules_ListByHealthModel.json
     */
    /**
     * Sample code: DiscoveryRules_ListByHealthModel.
     * 
     * @param manager Entry point to CloudHealthManager.
     */
    public static void
        discoveryRulesListByHealthModel(com.azure.resourcemanager.cloudhealth.CloudHealthManager manager) {
        manager.discoveryRules().listByHealthModel("online-store-rg", "online-store", null,
            com.azure.core.util.Context.NONE);
    }
}
