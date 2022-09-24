import com.azure.core.util.Context;
import com.azure.resourcemanager.monitor.fluent.models.LogSearchRuleResourceInner;
import com.azure.resourcemanager.monitor.models.AlertSeverity;
import com.azure.resourcemanager.monitor.models.AlertingAction;
import com.azure.resourcemanager.monitor.models.AzNsActionGroup;
import com.azure.resourcemanager.monitor.models.ConditionalOperator;
import com.azure.resourcemanager.monitor.models.Enabled;
import com.azure.resourcemanager.monitor.models.LogMetricTrigger;
import com.azure.resourcemanager.monitor.models.MetricTriggerType;
import com.azure.resourcemanager.monitor.models.QueryType;
import com.azure.resourcemanager.monitor.models.Schedule;
import com.azure.resourcemanager.monitor.models.Source;
import com.azure.resourcemanager.monitor.models.TriggerCondition;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for ScheduledQueryRules CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2018-04-16/examples/createOrUpdateScheduledQueryRules.json
     */
    /**
     * Sample code: Create or Update rule - AlertingAction.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateRuleAlertingAction(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .diagnosticSettings()
            .manager()
            .serviceClient()
            .getScheduledQueryRules()
            .createOrUpdateWithResponse(
                "Rac46PostSwapRG",
                "logalertfoo",
                new LogSearchRuleResourceInner()
                    .withLocation("eastus")
                    .withTags(mapOf())
                    .withDescription("log alert description")
                    .withEnabled(Enabled.TRUE)
                    .withSource(
                        new Source()
                            .withQuery("Heartbeat | summarize AggregatedValue = count() by bin(TimeGenerated, 5m)")
                            .withDataSourceId(
                                "/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/Rac46PostSwapRG/providers/Microsoft.OperationalInsights/workspaces/sampleWorkspace")
                            .withQueryType(QueryType.RESULT_COUNT))
                    .withSchedule(new Schedule().withFrequencyInMinutes(15).withTimeWindowInMinutes(15))
                    .withAction(
                        new AlertingAction()
                            .withSeverity(AlertSeverity.ONE)
                            .withAznsAction(
                                new AzNsActionGroup()
                                    .withActionGroup(Arrays.asList())
                                    .withEmailSubject("Email Header")
                                    .withCustomWebhookPayload("{}"))
                            .withTrigger(
                                new TriggerCondition()
                                    .withThresholdOperator(ConditionalOperator.GREATER_THAN)
                                    .withThreshold(3.0)
                                    .withMetricTrigger(
                                        new LogMetricTrigger()
                                            .withThresholdOperator(ConditionalOperator.GREATER_THAN)
                                            .withThreshold(5.0D)
                                            .withMetricTriggerType(MetricTriggerType.CONSECUTIVE)
                                            .withMetricColumn("Computer")))),
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
