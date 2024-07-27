
import com.azure.resourcemanager.monitor.fluent.models.MetricAlertResourceInner;
import com.azure.resourcemanager.monitor.models.WebtestLocationAvailabilityCriteria;
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
     * createOrUpdateWebTestMetricAlert.json
     */
    /**
     * Sample code: Create or update a web test alert rule.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateAWebTestAlertRule(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getMetricAlerts().createOrUpdateWithResponse("rg-example",
            "webtest-name-example",
            new MetricAlertResourceInner().withLocation("global").withTags(mapOf(
                "hidden-link:/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/components/webtest-name-example",
                "Resource",
                "hidden-link:/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/webtests/component-example",
                "Resource"))
                .withDescription("Automatically created alert rule for availability test \"component-example\" a")
                .withSeverity(4).withEnabled(true)
                .withScopes(Arrays.asList(
                    "/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/webtests/component-example",
                    "/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/components/webtest-name-example"))
                .withEvaluationFrequency(Duration.parse("PT1M")).withWindowSize(Duration.parse("PT15M"))
                .withCriteria(new WebtestLocationAvailabilityCriteria().withWebTestId(
                    "/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/webtests/component-example")
                    .withComponentId(
                        "/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/components/webtest-name-example")
                    .withFailedLocationCount(2f))
                .withActions(Arrays.asList()),
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
