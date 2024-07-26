
import com.azure.resourcemanager.monitor.fluent.models.MetricAlertResourceInner;
import com.azure.resourcemanager.monitor.models.AggregationTypeEnum;
import com.azure.resourcemanager.monitor.models.MetricAlertAction;
import com.azure.resourcemanager.monitor.models.MetricAlertSingleResourceMultipleMetricCriteria;
import com.azure.resourcemanager.monitor.models.MetricCriteria;
import com.azure.resourcemanager.monitor.models.Operator;
import java.time.Duration;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for MetricAlerts CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2018-03-01/examples/
     * createOrUpdateMetricAlertSingleResource.json
     */
    /**
     * Sample code: Create or update an alert rule for Single Resource.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createOrUpdateAnAlertRuleForSingleResource(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getMetricAlerts().createOrUpdateWithResponse("gigtest",
            "chiricutin",
            new MetricAlertResourceInner().withLocation("global").withTags(mapOf())
                .withDescription("This is the description of the rule1").withSeverity(3).withEnabled(true)
                .withScopes(Arrays.asList(
                    "/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourceGroups/gigtest/providers/Microsoft.Compute/virtualMachines/gigwadme"))
                .withEvaluationFrequency(Duration.parse("Pt1m")).withWindowSize(Duration.parse("Pt15m"))
                .withCriteria(new MetricAlertSingleResourceMultipleMetricCriteria().withAllOf(Arrays.asList(
                    new MetricCriteria().withName("High_CPU_80").withMetricName("\\Processor(_Total)\\% Processor Time")
                        .withTimeAggregation(AggregationTypeEnum.AVERAGE).withDimensions(Arrays.asList())
                        .withOperator(Operator.GREATER_THAN).withThreshold(80.5))))
                .withAutoMitigate(true)
                .withActions(Arrays.asList(new MetricAlertAction().withActionGroupId(
                    "/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourcegroups/gigtest/providers/microsoft.insights/actiongroups/group2")
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
