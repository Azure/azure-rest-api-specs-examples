
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
     * 2026-04-01/AdvancedPlatformMetricsCRUD/AdvancedPlatformMetricsRules_CreateOrUpdate_ContainerList.json
     */
    /**
     * Sample code: AdvancedPlatformMetricsRules_CreateOrUpdate_ContainerList - Create advanced platform metrics rule
     * with container list filter.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void
        advancedPlatformMetricsRulesCreateOrUpdateContainerListCreateAdvancedPlatformMetricsRuleWithContainerListFilter(
            com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getAdvancedPlatformMetrics().createOrUpdateWithResponse("res6977", "sto2527",
            AdvancedPlatformMetricsRuleType.CONTAINER_LEVEL_CAPACITY_METRICS,
            new AdvancedPlatformMetricsRuleInner()
                .withProperties(new AdvancedPlatformMetricsRuleProperties().withEnabled(true)
                    .withRuleConfig(new AdvancedPlatformMetricsRuleConfig()
                        .withFilterType(AdvancedPlatformMetricsFilterType.CONTAINER_LIST_FILTER)
                        .withFilterValues(Arrays.asList("container1", "container2", "container3")))),
            com.azure.core.util.Context.NONE);
    }
}
