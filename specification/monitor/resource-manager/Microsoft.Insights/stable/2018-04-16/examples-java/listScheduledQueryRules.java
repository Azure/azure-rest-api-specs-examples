import com.azure.core.util.Context;

/** Samples for ScheduledQueryRules ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2018-04-16/examples/listScheduledQueryRules.json
     */
    /**
     * Sample code: List rules.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listRules(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .diagnosticSettings()
            .manager()
            .serviceClient()
            .getScheduledQueryRules()
            .listByResourceGroup("gigtest", null, Context.NONE);
    }
}
