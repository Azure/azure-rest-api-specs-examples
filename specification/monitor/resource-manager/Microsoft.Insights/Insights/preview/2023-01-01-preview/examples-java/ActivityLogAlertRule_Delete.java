
/**
 * Samples for ActivityLogAlerts Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-01-01-preview/ActivityLogAlertRule_Delete.json
     */
    /**
     * Sample code: Delete an Activity Log Alert rule.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void deleteAnActivityLogAlertRule(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getActivityLogAlerts().deleteWithResponse("MyResourceGroup",
            "SampleActivityLogAlertRule", com.azure.core.util.Context.NONE);
    }
}
