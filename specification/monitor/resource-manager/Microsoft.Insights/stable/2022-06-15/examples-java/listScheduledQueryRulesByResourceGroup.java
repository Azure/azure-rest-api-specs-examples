
/**
 * Samples for ScheduledQueryRules ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2022-06-15/examples/
     * listScheduledQueryRulesByResourceGroup.json
     */
    /**
     * Sample code: List scheduled query rules by resource group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listScheduledQueryRulesByResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getScheduledQueryRules()
            .listByResourceGroup("QueryResourceGroupName", com.azure.core.util.Context.NONE);
    }
}
