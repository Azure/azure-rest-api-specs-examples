
/**
 * Samples for MetricAlerts Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/deleteMetricAlert.json
     */
    /**
     * Sample code: Delete an alert rule.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void deleteAnAlertRule(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getMetricAlerts().deleteWithResponse("gigtest", "chiricutin",
            com.azure.core.util.Context.NONE);
    }
}
