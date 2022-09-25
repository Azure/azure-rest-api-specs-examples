import com.azure.core.util.Context;

/** Samples for ScheduledQueryRules GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2018-04-16/examples/getScheduledQueryRules.json
     */
    /**
     * Sample code: Get rule.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getRule(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .diagnosticSettings()
            .manager()
            .serviceClient()
            .getScheduledQueryRules()
            .getByResourceGroupWithResponse("Rac46PostSwapRG", "logalertfoo", Context.NONE);
    }
}
