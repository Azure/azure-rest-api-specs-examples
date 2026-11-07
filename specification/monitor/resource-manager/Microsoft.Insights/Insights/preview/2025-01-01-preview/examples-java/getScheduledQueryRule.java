
/**
 * Samples for ScheduledQueryRules GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01-preview/getScheduledQueryRule.json
     */
    /**
     * Sample code: Get a scheduled query rule for single resource.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void
        getAScheduledQueryRuleForSingleResource(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getScheduledQueryRules().getByResourceGroupWithResponse("QueryResourceGroupName",
            "perf", com.azure.core.util.Context.NONE);
    }
}
