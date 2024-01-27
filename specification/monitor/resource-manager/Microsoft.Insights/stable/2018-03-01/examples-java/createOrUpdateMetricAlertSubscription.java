
import com.azure.core.util.Context;
import com.azure.resourcemanager.monitor.fluent.models.MetricAlertResourceInner;
import com.azure.resourcemanager.monitor.models.AggregationTypeEnum;
import com.azure.resourcemanager.monitor.models.MetricAlertAction;
import com.azure.resourcemanager.monitor.models.MetricAlertMultipleResourceMultipleMetricCriteria;
import com.azure.resourcemanager.monitor.models.MetricCriteria;
import com.azure.resourcemanager.monitor.models.Operator;
import java.time.Duration;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for MetricAlerts CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2018-03-01/examples/
     * createOrUpdateMetricAlertSubscription.json
     */
    /**
     * Sample code: Create or update an alert rule on Subscription.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateAnAlertRuleOnSubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getMetricAlerts().createOrUpdateWithResponse("gigtest",
            "MetricAlertAtSubscriptionLevel",
            new MetricAlertResourceInner().withLocation("global").withTags(mapOf())
                .withDescription("This is the description of the rule1").withSeverity(3).withEnabled(true)
                .withScopes(Arrays.asList("/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7"))
                .withEvaluationFrequency(Duration.parse("PT1M")).withWindowSize(Duration.parse("PT15M"))
                .withTargetResourceType("Microsoft.Compute/virtualMachines").withTargetResourceRegion("southcentralus")
                .withCriteria(new MetricAlertMultipleResourceMultipleMetricCriteria()
                    .withAllOf(Arrays.asList(new MetricCriteria().withName("High_CPU_80")
                        .withMetricName("Percentage CPU").withMetricNamespace("microsoft.compute/virtualmachines")
                        .withTimeAggregation(AggregationTypeEnum.AVERAGE).withDimensions(Arrays.asList())
                        .withOperator(Operator.GREATER_THAN).withThreshold(80.5))))
                .withAutoMitigate(true)
                .withActions(Arrays.asList(new MetricAlertAction().withActionGroupId(
                    "/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourcegroups/gigtest/providers/microsoft.insights/actiongroups/group2")
                    .withWebhookProperties(mapOf("key11", "value11", "key12", "value12")))),
            Context.NONE);
    }

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
