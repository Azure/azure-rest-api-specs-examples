
/**
 * Samples for ScheduledQueryRules List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01-preview/listScheduledQueryRulesBySubscription.json
     */
    /**
     * Sample code: List scheduled query rules by subscription.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void listScheduledQueryRulesBySubscription(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getScheduledQueryRules().list(com.azure.core.util.Context.NONE);
    }
}
