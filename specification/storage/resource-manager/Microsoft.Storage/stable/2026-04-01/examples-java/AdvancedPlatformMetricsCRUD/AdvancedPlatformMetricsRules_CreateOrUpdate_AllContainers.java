
import com.azure.resourcemanager.storage.fluent.models.AdvancedPlatformMetricsRuleInner;
import com.azure.resourcemanager.storage.models.AdvancedPlatformMetricsFilterType;
import com.azure.resourcemanager.storage.models.AdvancedPlatformMetricsRuleConfig;
import com.azure.resourcemanager.storage.models.AdvancedPlatformMetricsRuleProperties;
import com.azure.resourcemanager.storage.models.AdvancedPlatformMetricsRuleType;

/**
 * Samples for AdvancedPlatformMetrics CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2026-04-01/AdvancedPlatformMetricsCRUD/AdvancedPlatformMetricsRules_CreateOrUpdate_AllContainers.json
     */
    /**
     * Sample code: AdvancedPlatformMetricsRules_CreateOrUpdate_AllContainers - Create advanced platform metrics rule
     * with all containers filter.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void
        advancedPlatformMetricsRulesCreateOrUpdateAllContainersCreateAdvancedPlatformMetricsRuleWithAllContainersFilter(
            com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getAdvancedPlatformMetrics().createOrUpdateWithResponse("res6977", "sto2527",
            AdvancedPlatformMetricsRuleType.CONTAINER_LEVEL_CAPACITY_METRICS,
            new AdvancedPlatformMetricsRuleInner()
                .withProperties(new AdvancedPlatformMetricsRuleProperties().withEnabled(true)
                    .withRuleConfig(new AdvancedPlatformMetricsRuleConfig()
                        .withFilterType(AdvancedPlatformMetricsFilterType.ALL_CONTAINERS_FILTER))),
            com.azure.core.util.Context.NONE);
    }
}
