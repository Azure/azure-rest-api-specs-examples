import com.azure.core.util.Context;

/** Samples for ScheduledQueryRules Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2018-04-16/examples/deleteScheduledQueryRules.json
     */
    /**
     * Sample code: Delete rule.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteRule(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .diagnosticSettings()
            .manager()
            .serviceClient()
            .getScheduledQueryRules()
            .deleteWithResponse("Rac46PostSwapRG", "logalertfoo", Context.NONE);
    }
}
