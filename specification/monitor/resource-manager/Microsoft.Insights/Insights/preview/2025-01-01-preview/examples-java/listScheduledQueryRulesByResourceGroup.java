
/**
 * Samples for ScheduledQueryRules ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01-preview/listScheduledQueryRulesByResourceGroup.json
     */
    /**
     * Sample code: List scheduled query rules by resource group.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void
        listScheduledQueryRulesByResourceGroup(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getScheduledQueryRules().listByResourceGroup("QueryResourceGroupName",
            com.azure.core.util.Context.NONE);
    }
}
