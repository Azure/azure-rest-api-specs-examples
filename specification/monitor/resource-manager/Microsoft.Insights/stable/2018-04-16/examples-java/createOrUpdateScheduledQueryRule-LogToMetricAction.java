import com.azure.core.util.Context;
import com.azure.resourcemanager.monitor.fluent.models.LogSearchRuleResourceInner;
import com.azure.resourcemanager.monitor.models.Criteria;
import com.azure.resourcemanager.monitor.models.Enabled;
import com.azure.resourcemanager.monitor.models.LogToMetricAction;
import com.azure.resourcemanager.monitor.models.Source;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for ScheduledQueryRules CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2018-04-16/examples/createOrUpdateScheduledQueryRule-LogToMetricAction.json
     */
    /**
     * Sample code: Create or Update rule - LogToMetricAction.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateRuleLogToMetricAction(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .diagnosticSettings()
            .manager()
            .serviceClient()
            .getScheduledQueryRules()
            .createOrUpdateWithResponse(
                "alertsweu",
                "logtometricfoo",
                new LogSearchRuleResourceInner()
                    .withLocation("West Europe")
                    .withTags(mapOf())
                    .withDescription("log to metric description")
                    .withEnabled(Enabled.TRUE)
                    .withSource(
                        new Source()
                            .withDataSourceId(
                                "/subscriptions/af52d502-a447-4bc6-8cb7-4780fbb00490/resourceGroups/alertsweu/providers/Microsoft.OperationalInsights/workspaces/alertsweu"))
                    .withAction(
                        new LogToMetricAction()
                            .withCriteria(
                                Arrays
                                    .asList(
                                        new Criteria()
                                            .withMetricName("Average_% Idle Time")
                                            .withDimensions(Arrays.asList())))),
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
