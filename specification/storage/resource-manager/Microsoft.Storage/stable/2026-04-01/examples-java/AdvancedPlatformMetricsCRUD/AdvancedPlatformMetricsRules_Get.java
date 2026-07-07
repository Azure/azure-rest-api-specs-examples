
import com.azure.resourcemanager.storage.models.AdvancedPlatformMetricsRuleType;

/**
 * Samples for AdvancedPlatformMetrics Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/AdvancedPlatformMetricsCRUD/AdvancedPlatformMetricsRules_Get.json
     */
    /**
     * Sample code: AdvancedPlatformMetricsRules_Get - Get advanced platform metrics rule.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void advancedPlatformMetricsRulesGetGetAdvancedPlatformMetricsRule(
        com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getAdvancedPlatformMetrics().getWithResponse("res6977", "sto2527",
            AdvancedPlatformMetricsRuleType.CONTAINER_LEVEL_CAPACITY_METRICS, com.azure.core.util.Context.NONE);
    }
}
