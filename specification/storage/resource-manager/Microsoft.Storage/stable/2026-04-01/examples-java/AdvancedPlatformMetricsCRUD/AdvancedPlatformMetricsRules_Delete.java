
import com.azure.resourcemanager.storage.models.AdvancedPlatformMetricsRuleType;

/**
 * Samples for AdvancedPlatformMetrics Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/AdvancedPlatformMetricsCRUD/AdvancedPlatformMetricsRules_Delete.json
     */
    /**
     * Sample code: AdvancedPlatformMetricsRules_Delete - Delete advanced platform metrics rule.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void advancedPlatformMetricsRulesDeleteDeleteAdvancedPlatformMetricsRule(
        com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getAdvancedPlatformMetrics().deleteWithResponse("res6977", "sto2527",
            AdvancedPlatformMetricsRuleType.CONTAINER_LEVEL_CAPACITY_METRICS, com.azure.core.util.Context.NONE);
    }
}
