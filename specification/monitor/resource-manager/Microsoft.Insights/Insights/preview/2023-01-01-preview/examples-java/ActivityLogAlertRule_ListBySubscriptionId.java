
/**
 * Samples for ActivityLogAlerts List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-01-01-preview/ActivityLogAlertRule_ListBySubscriptionId.json
     */
    /**
     * Sample code: Get list of all Activity Log Alert rules under a subscription.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void
        getListOfAllActivityLogAlertRulesUnderASubscription(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getActivityLogAlerts().list(com.azure.core.util.Context.NONE);
    }
}
