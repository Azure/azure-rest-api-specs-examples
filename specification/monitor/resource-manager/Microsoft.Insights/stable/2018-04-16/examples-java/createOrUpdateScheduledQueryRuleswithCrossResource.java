import com.azure.core.util.Context;
import com.azure.resourcemanager.monitor.fluent.models.LogSearchRuleResourceInner;
import com.azure.resourcemanager.monitor.models.AlertSeverity;
import com.azure.resourcemanager.monitor.models.AlertingAction;
import com.azure.resourcemanager.monitor.models.AzNsActionGroup;
import com.azure.resourcemanager.monitor.models.ConditionalOperator;
import com.azure.resourcemanager.monitor.models.Enabled;
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
     * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2018-04-16/examples/createOrUpdateScheduledQueryRuleswithCrossResource.json
     */
    /**
     * Sample code: Create or Update rule - AlertingAction with Cross-Resource.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateRuleAlertingActionWithCrossResource(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .diagnosticSettings()
            .manager()
            .serviceClient()
            .getScheduledQueryRules()
            .createOrUpdateWithResponse(
                "Rac46PostSwapRG",
                "SampleCrossResourceAlert",
                new LogSearchRuleResourceInner()
                    .withLocation("eastus")
                    .withTags(mapOf())
                    .withDescription("Sample Cross Resource alert")
                    .withEnabled(Enabled.TRUE)
                    .withSource(
                        new Source()
                            .withQuery("union requests, workspace(\"sampleWorkspace\").Update")
                            .withAuthorizedResources(
                                Arrays
                                    .asList(
                                        "/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/Rac46PostSwapRG/providers/Microsoft.OperationalInsights/workspaces/sampleWorkspace",
                                        "/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/Rac46PostSwapRG/providers/microsoft.insights/components/sampleAI"))
                            .withDataSourceId(
                                "/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/Rac46PostSwapRG/providers/microsoft.insights/components/sampleAI")
                            .withQueryType(QueryType.RESULT_COUNT))
                    .withSchedule(new Schedule().withFrequencyInMinutes(60).withTimeWindowInMinutes(60))
                    .withAction(
                        new AlertingAction()
                            .withSeverity(AlertSeverity.THREE)
                            .withAznsAction(
                                new AzNsActionGroup()
                                    .withActionGroup(
                                        Arrays
                                            .asList(
                                                "/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/Rac46PostSwapRG/providers/microsoft.insights/actiongroups/test-ag"))
                                    .withEmailSubject("Cross Resource Mail!!"))
                            .withTrigger(
                                new TriggerCondition()
                                    .withThresholdOperator(ConditionalOperator.GREATER_THAN)
                                    .withThreshold(5000.0))),
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
