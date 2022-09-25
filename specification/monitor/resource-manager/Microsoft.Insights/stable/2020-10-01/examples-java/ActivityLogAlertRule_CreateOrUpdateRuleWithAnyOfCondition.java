import com.azure.core.util.Context;
import com.azure.resourcemanager.monitor.fluent.models.ActivityLogAlertResourceInner;
import com.azure.resourcemanager.monitor.models.ActionList;
import com.azure.resourcemanager.monitor.models.ActivityLogAlertActionGroup;
import com.azure.resourcemanager.monitor.models.ActivityLogAlertAllOfCondition;
import com.azure.resourcemanager.monitor.models.ActivityLogAlertLeafCondition;
import com.azure.resourcemanager.monitor.models.AlertRuleLeafCondition;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for ActivityLogAlerts CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2020-10-01/examples/ActivityLogAlertRule_CreateOrUpdateRuleWithAnyOfCondition.json
     */
    /**
     * Sample code: Create or update an Activity Log Alert rule with 'anyOf' condition.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateAnActivityLogAlertRuleWithAnyOfCondition(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .diagnosticSettings()
            .manager()
            .serviceClient()
            .getActivityLogAlerts()
            .createOrUpdateWithResponse(
                "MyResourceGroup",
                "SampleActivityLogAlertRuleWithAnyOfCondition",
                new ActivityLogAlertResourceInner()
                    .withLocation("Global")
                    .withTags(mapOf())
                    .withScopes(Arrays.asList("subscriptions/187f412d-1758-44d9-b052-169e2564721d"))
                    .withCondition(
                        new ActivityLogAlertAllOfCondition()
                            .withAllOf(
                                Arrays
                                    .asList(
                                        new ActivityLogAlertLeafCondition()
                                            .withField("category")
                                            .withEquals("ServiceHealth"),
                                        new ActivityLogAlertLeafCondition()
                                            .withAnyOf(
                                                Arrays
                                                    .asList(
                                                        new AlertRuleLeafCondition()
                                                            .withField("properties.incidentType")
                                                            .withEquals("Incident"),
                                                        new AlertRuleLeafCondition()
                                                            .withField("properties.incidentType")
                                                            .withEquals("Maintenance"))))))
                    .withActions(
                        new ActionList()
                            .withActionGroups(
                                Arrays
                                    .asList(
                                        new ActivityLogAlertActionGroup()
                                            .withActionGroupId(
                                                "/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/MyResourceGroup/providers/Microsoft.Insights/actionGroups/SampleActionGroup")
                                            .withWebhookProperties(
                                                mapOf("sampleWebhookProperty", "SamplePropertyValue")))))
                    .withEnabled(true)
                    .withDescription("Description of sample Activity Log Alert rule with 'anyOf' condition."),
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
