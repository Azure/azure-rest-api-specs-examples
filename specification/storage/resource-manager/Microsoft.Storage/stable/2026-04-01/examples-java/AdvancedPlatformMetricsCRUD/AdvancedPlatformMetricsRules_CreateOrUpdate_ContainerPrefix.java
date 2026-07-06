
import com.azure.resourcemanager.storage.fluent.models.AdvancedPlatformMetricsRuleInner;
import com.azure.resourcemanager.storage.models.AdvancedPlatformMetricsFilterType;
import com.azure.resourcemanager.storage.models.AdvancedPlatformMetricsRuleConfig;
import com.azure.resourcemanager.storage.models.AdvancedPlatformMetricsRuleProperties;
import com.azure.resourcemanager.storage.models.AdvancedPlatformMetricsRuleType;
import java.util.Arrays;

/**
 * Samples for AdvancedPlatformMetrics CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2026-04-01/AdvancedPlatformMetricsCRUD/AdvancedPlatformMetricsRules_CreateOrUpdate_ContainerPrefix.json
     */
    /**
     * Sample code: AdvancedPlatformMetricsRules_CreateOrUpdate_ContainerPrefix - Create advanced platform metrics rule
     * with container prefix filter.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void
        advancedPlatformMetricsRulesCreateOrUpdateContainerPrefixCreateAdvancedPlatformMetricsRuleWithContainerPrefixFilter(
            com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getAdvancedPlatformMetrics().createOrUpdateWithResponse("res6977", "sto2527",
            AdvancedPlatformMetricsRuleType.CONTAINER_LEVEL_CAPACITY_METRICS,
            new AdvancedPlatformMetricsRuleInner()
                .withProperties(new AdvancedPlatformMetricsRuleProperties().withEnabled(true)
                    .withRuleConfig(new AdvancedPlatformMetricsRuleConfig()
                        .withFilterType(AdvancedPlatformMetricsFilterType.CONTAINER_PREFIX_FILTER)
                        .withFilterValues(Arrays.asList("logs", "data")))),
            com.azure.core.util.Context.NONE);
    }
}
