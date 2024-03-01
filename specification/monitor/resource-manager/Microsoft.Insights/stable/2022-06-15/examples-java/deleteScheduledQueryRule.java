
/**
 * Samples for ScheduledQueryRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/stable/2022-06-15/examples/deleteScheduledQueryRule.
     * json
     */
    /**
     * Sample code: Delete a scheduled query rule.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAScheduledQueryRule(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getScheduledQueryRules()
            .deleteWithResponse("QueryResourceGroupName", "heartbeat", com.azure.core.util.Context.NONE);
    }
}
