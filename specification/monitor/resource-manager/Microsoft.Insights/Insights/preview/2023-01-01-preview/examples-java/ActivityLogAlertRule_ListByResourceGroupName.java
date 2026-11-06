
/**
 * Samples for ActivityLogAlerts ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-01-01-preview/ActivityLogAlertRule_ListByResourceGroupName.json
     */
    /**
     * Sample code: List activity log alerts.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void listActivityLogAlerts(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getActivityLogAlerts().listByResourceGroup("MyResourceGroup",
            com.azure.core.util.Context.NONE);
    }
}
