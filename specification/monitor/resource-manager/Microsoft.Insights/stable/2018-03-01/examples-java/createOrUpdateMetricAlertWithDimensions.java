
import com.azure.core.util.Context;
import com.azure.resourcemanager.monitor.fluent.models.MetricAlertResourceInner;
import com.azure.resourcemanager.monitor.models.AggregationTypeEnum;
import com.azure.resourcemanager.monitor.models.MetricAlertAction;
import com.azure.resourcemanager.monitor.models.MetricAlertMultipleResourceMultipleMetricCriteria;
import com.azure.resourcemanager.monitor.models.MetricCriteria;
import com.azure.resourcemanager.monitor.models.MetricDimension;
import com.azure.resourcemanager.monitor.models.Operator;
import java.time.Duration;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for MetricAlerts CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2018-03-01/examples/
     * createOrUpdateMetricAlertWithDimensions.json
     */
    /**
     * Sample code: Create or update an alert rules with dimensions.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateAnAlertRulesWithDimensions(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getMetricAlerts().createOrUpdateWithResponse("gigtest",
            "MetricAlertOnMultipleDimensions",
            new MetricAlertResourceInner().withLocation("global").withTags(mapOf())
                .withDescription("This is the description of the rule1").withSeverity(3).withEnabled(true)
                .withScopes(Arrays.asList(
                    "/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourceGroups/gigtest/providers/Microsoft.KeyVault/vaults/keyVaultResource"))
                .withEvaluationFrequency(Duration.parse("PT1H")).withWindowSize(Duration.parse("P1D"))
                .withCriteria(new MetricAlertMultipleResourceMultipleMetricCriteria()
                    .withAllOf(Arrays.asList(new MetricCriteria().withName("Metric1").withMetricName("Availability")
                        .withMetricNamespace("Microsoft.KeyVault/vaults")
                        .withTimeAggregation(AggregationTypeEnum.AVERAGE)
                        .withDimensions(Arrays.asList(
                            new MetricDimension().withName("ActivityName").withOperator("Include")
                                .withValues(Arrays.asList("*")),
                            new MetricDimension().withName("StatusCode").withOperator("Include")
                                .withValues(Arrays.asList("200"))))
                        .withOperator(Operator.GREATER_THAN).withThreshold(55.0))))
                .withAutoMitigate(true)
                .withActions(Arrays.asList(new MetricAlertAction().withActionGroupId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/gigtest/providers/microsoft.insights/actiongroups/group2")
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
