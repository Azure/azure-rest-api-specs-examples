import com.azure.core.util.Context;
import com.azure.resourcemanager.monitor.models.Enabled;
import com.azure.resourcemanager.monitor.models.LogSearchRuleResourcePatch;

/** Samples for ScheduledQueryRules Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2018-04-16/examples/patchScheduledQueryRules.json
     */
    /**
     * Sample code: Patch Log Search Rule.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void patchLogSearchRule(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .diagnosticSettings()
            .manager()
            .serviceClient()
            .getScheduledQueryRules()
            .updateWithResponse(
                "my-resource-group",
                "logalertfoo",
                new LogSearchRuleResourcePatch().withEnabled(Enabled.TRUE),
                Context.NONE);
    }
}
