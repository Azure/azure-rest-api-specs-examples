
import com.azure.resourcemanager.monitor.fluent.models.ScheduledQueryRuleResourceInner;
import com.azure.resourcemanager.monitor.models.Actions;
import com.azure.resourcemanager.monitor.models.AlertSeverity;
import com.azure.resourcemanager.monitor.models.Condition;
import com.azure.resourcemanager.monitor.models.ConditionOperator;
import com.azure.resourcemanager.monitor.models.CriterionType;
import com.azure.resourcemanager.monitor.models.Dimension;
import com.azure.resourcemanager.monitor.models.DimensionOperator;
import com.azure.resourcemanager.monitor.models.ScheduledQueryRuleCriteria;
import com.azure.resourcemanager.monitor.models.TimeAggregation;
import java.time.Duration;
import java.time.OffsetDateTime;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ScheduledQueryRules CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01-preview/createOrUpdateDynamicThresholdScheduledQueryRule.json
     */
    /**
     * Sample code: create or update dynamic threshold scheduled query rule on Subscription.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void createOrUpdateDynamicThresholdScheduledQueryRuleOnSubscription(
        com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getScheduledQueryRules().createOrUpdateWithResponse("QueryResourceGroupName", "perf",
            new ScheduledQueryRuleResourceInner().withLocation("eastus").withDescription("Performance rule")
                .withSeverity(AlertSeverity.FOUR).withEnabled(true)
                .withScopes(Arrays.asList(
                    "/subscriptions/aaf177ed-1330-a9f2-80ea-fd3d7783b147/resourceGroups/scopeResourceGroup1/providers/Microsoft.Compute/virtualMachines/vm1"))
                .withEvaluationFrequency(Duration.parse("PT5M")).withWindowSize(Duration.parse("PT10M"))
                .withCriteria(new ScheduledQueryRuleCriteria().withAllOf(Arrays.asList(new Condition()
                    .withCriterionType(CriterionType.DYNAMIC_THRESHOLD_CRITERION)
                    .withQuery("Perf | where ObjectName == \"Processor\"").withTimeAggregation(TimeAggregation.AVERAGE)
                    .withMetricMeasureColumn("% Processor Time").withResourceIdColumn("resourceId")
                    .withDimensions(Arrays.asList(
                        new Dimension().withName("ComputerIp").withOperator(DimensionOperator.EXCLUDE)
                            .withValues(Arrays.asList("192.168.1.1")),
                        new Dimension().withName("OSType").withOperator(DimensionOperator.INCLUDE)
                            .withValues(Arrays.asList("*"))))
                    .withOperator(ConditionOperator.GREATER_OR_LESS_THAN).withAlertSensitivity("Medium")
                    .withIgnoreDataBefore(OffsetDateTime.parse("2024-06-01T21:00:00.000Z")))))
                .withMuteActionsDuration(Duration.parse("PT30M"))
                .withActions(new Actions().withActionGroups(Arrays.asList(
                    "/subscriptions/1cf177ed-1330-4692-80ea-fd3d7783b147/resourcegroups/sqrapi/providers/microsoft.insights/actiongroups/myactiongroup"))
                    .withCustomProperties(mapOf("key11", "fakeTokenPlaceholder", "key12", "fakeTokenPlaceholder"))
                    .withActionProperties(mapOf("Icm.Title", "Custom title in ICM", "Icm.TsgId", "https://tsg.url")))
                .withCheckWorkspaceAlertsStorageConfigured(true).withSkipQueryValidation(true),
            com.azure.core.util.Context.NONE);
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
