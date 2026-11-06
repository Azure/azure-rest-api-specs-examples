
/**
 * Samples for ScheduledQueryRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01-preview/deleteScheduledQueryRule.json
     */
    /**
     * Sample code: Delete a scheduled query rule.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void deleteAScheduledQueryRule(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getScheduledQueryRules().deleteWithResponse("QueryResourceGroupName", "heartbeat",
            com.azure.core.util.Context.NONE);
    }
}
