
import com.azure.resourcemanager.monitor.fluent.models.ScheduledQueryRuleResourceInner;
import com.azure.resourcemanager.monitor.models.Actions;
import com.azure.resourcemanager.monitor.models.AlertSeverity;
import com.azure.resourcemanager.monitor.models.Condition;
import com.azure.resourcemanager.monitor.models.Kind;
import com.azure.resourcemanager.monitor.models.ScheduledQueryRuleCriteria;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ScheduledQueryRules CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01-preview/createOrUpdateSimpleLogAlertScheduledQueryRule.json
     */
    /**
     * Sample code: Create or update a simple log alert scheduled query rule on Subscription.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void createOrUpdateASimpleLogAlertScheduledQueryRuleOnSubscription(
        com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getScheduledQueryRules().createOrUpdateWithResponse("QueryResourceGroupName", "perf",
            new ScheduledQueryRuleResourceInner().withLocation("eastus").withKind(Kind.SIMPLE_LOG_ALERT)
                .withDescription("Performance rule").withSeverity(AlertSeverity.FOUR).withEnabled(true)
                .withScopes(Arrays.asList(
                    "/subscriptions/aaf177ed-1330-a9f2-80ea-fd3d7783b147/resourceGroups/scopeResourceGroup1/providers/Microsoft.Compute/virtualMachines/vm1"))
                .withCriteria(new ScheduledQueryRuleCriteria()
                    .withAllOf(Arrays.asList(new Condition().withQuery("Perf | where ObjectName == \"Processor\""))))
                .withActions(new Actions().withActionGroups(Arrays.asList(
                    "/subscriptions/1cf177ed-1330-4692-80ea-fd3d7783b147/resourcegroups/sqrapi/providers/microsoft.insights/actiongroups/myactiongroup"))
                    .withCustomProperties(mapOf("key11", "fakeTokenPlaceholder", "key12", "fakeTokenPlaceholder"))
                    .withActionProperties(mapOf("Icm.Title", "Custom title in ICM", "Icm.TsgId", "https://tsg.url")))
                .withCheckWorkspaceAlertsStorageConfigured(true).withSkipQueryValidation(true).withAutoMitigate(false),
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
