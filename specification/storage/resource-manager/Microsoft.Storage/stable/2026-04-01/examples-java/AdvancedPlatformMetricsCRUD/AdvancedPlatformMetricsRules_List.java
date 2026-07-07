
/**
 * Samples for AdvancedPlatformMetrics List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/AdvancedPlatformMetricsCRUD/AdvancedPlatformMetricsRules_List.json
     */
    /**
     * Sample code: AdvancedPlatformMetricsRules_List - List advanced platform metrics rules.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void advancedPlatformMetricsRulesListListAdvancedPlatformMetricsRules(
        com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getAdvancedPlatformMetrics().list("res6977", "sto2527",
            com.azure.core.util.Context.NONE);
    }
}
