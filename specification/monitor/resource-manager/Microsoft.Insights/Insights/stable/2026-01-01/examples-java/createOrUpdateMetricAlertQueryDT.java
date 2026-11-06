
import com.azure.resourcemanager.monitor.fluent.models.MetricAlertResourceInner;
import com.azure.resourcemanager.monitor.models.DynamicPromQLCriteria;
import com.azure.resourcemanager.monitor.models.DynamicThresholdOperator;
import com.azure.resourcemanager.monitor.models.DynamicThresholdSensitivity;
import com.azure.resourcemanager.monitor.models.Identity;
import com.azure.resourcemanager.monitor.models.IdentityType;
import com.azure.resourcemanager.monitor.models.MetricAlertAction;
import com.azure.resourcemanager.monitor.models.PromQLCriteria;
import com.azure.resourcemanager.monitor.models.QueryFailingPeriods;
import com.azure.resourcemanager.monitor.models.ResolveConfiguration;
import com.azure.resourcemanager.monitor.models.UserIdentityProperties;
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
     * x-ms-original-file: 2026-01-01/createOrUpdateMetricAlertQueryDT.json
     */
    /**
     * Sample code: Create or update a query based alert rule with dynamic threshold.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void createOrUpdateAQueryBasedAlertRuleWithDynamicThreshold(
        com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getMetricAlerts().createOrUpdateWithResponse("gigtest", "chiricutin",
            new MetricAlertResourceInner().withLocation("eastus").withTags(mapOf())
                .withIdentity(new Identity().withType(IdentityType.USER_ASSIGNED).withUserAssignedIdentities(mapOf(
                    "/subscriptions/2f1a501a-6e1d-4f37-a445-462d7f8a563d/resourceGroups/AdisTest/providers/Microsoft.ManagedIdentity/userAssignedIdentities/msi-test-euap",
                    new UserIdentityProperties())))
                .withDescription("This is the description of the rule1").withSeverity(3).withEnabled(true)
                .withScopes(Arrays.asList(
                    "/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourceGroups/gigtest/providers/microsoft.monitor/accounts/gigwadme"))
                .withEvaluationFrequency(Duration.parse("PT1M"))
                .withCriteria(new PromQLCriteria()
                    .withFailingPeriods(new QueryFailingPeriods().withForProperty(Duration.parse("PT5M")))
                    .withAllOf(Arrays.asList(new DynamicPromQLCriteria().withName("Metric1")
                        .withQuery("avg({\"system.cpu.utilization\"})").withOperator(DynamicThresholdOperator.LESS_THAN)
                        .withAlertSensitivity(DynamicThresholdSensitivity.MEDIUM)
                        .withIgnoreDataBefore(OffsetDateTime.parse("2019-04-04T21:00:00.000Z")))))
                .withResolveConfiguration(
                    new ResolveConfiguration().withAutoResolved(true).withTimeToResolve(Duration.parse("PT10M")))
                .withActions(Arrays.asList(new MetricAlertAction().withActionGroupId(
                    "/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourcegroups/gigtest/providers/microsoft.insights/actiongroups/group2")))
                .withCustomProperties(mapOf("key11", "fakeTokenPlaceholder", "key12", "fakeTokenPlaceholder"))
                .withActionProperties(mapOf("Email.Sujbect", "my custom email subject")),
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
