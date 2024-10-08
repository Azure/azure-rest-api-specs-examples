
import com.azure.resourcemanager.monitor.fluent.models.MetricAlertResourceInner;
import com.azure.resourcemanager.monitor.models.AggregationTypeEnum;
import com.azure.resourcemanager.monitor.models.DynamicMetricCriteria;
import com.azure.resourcemanager.monitor.models.DynamicThresholdFailingPeriods;
import com.azure.resourcemanager.monitor.models.DynamicThresholdOperator;
import com.azure.resourcemanager.monitor.models.DynamicThresholdSensitivity;
import com.azure.resourcemanager.monitor.models.MetricAlertAction;
import com.azure.resourcemanager.monitor.models.MetricAlertMultipleResourceMultipleMetricCriteria;
import java.time.Duration;
import java.time.OffsetDateTime;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for MetricAlerts CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2018-03-01/examples/
     * createOrUpdateDynamicMetricAlertSingleResource.json
     */
    /**
     * Sample code: Create or update a dynamic alert rule for Single Resource.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createOrUpdateADynamicAlertRuleForSingleResource(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getMetricAlerts().createOrUpdateWithResponse("gigtest",
            "chiricutin",
            new MetricAlertResourceInner().withLocation("global").withTags(mapOf())
                .withDescription("This is the description of the rule1").withSeverity(3).withEnabled(true)
                .withScopes(Arrays.asList(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/gigtest/providers/Microsoft.Compute/virtualMachines/gigwadme"))
                .withEvaluationFrequency(Duration.parse("PT1M")).withWindowSize(Duration.parse("PT15M"))
                .withCriteria(new MetricAlertMultipleResourceMultipleMetricCriteria()
                    .withAllOf(Arrays.asList(new DynamicMetricCriteria().withName("High_CPU_80")
                        .withMetricName("Percentage CPU").withMetricNamespace("microsoft.compute/virtualmachines")
                        .withTimeAggregation(AggregationTypeEnum.AVERAGE).withDimensions(Arrays.asList())
                        .withOperator(DynamicThresholdOperator.GREATER_OR_LESS_THAN)
                        .withAlertSensitivity(DynamicThresholdSensitivity.MEDIUM)
                        .withFailingPeriods(new DynamicThresholdFailingPeriods().withNumberOfEvaluationPeriods(4f)
                            .withMinFailingPeriodsToAlert(4f))
                        .withIgnoreDataBefore(OffsetDateTime.parse("2019-04-04T21:00:00.000Z")))))
                .withAutoMitigate(true)
                .withActions(Arrays.asList(new MetricAlertAction().withActionGroupId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/gigtest/providers/microsoft.insights/actiongroups/group2")
                    .withWebhookProperties(mapOf("key11", "fakeTokenPlaceholder", "key12", "fakeTokenPlaceholder")))),
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
