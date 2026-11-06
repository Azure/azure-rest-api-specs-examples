
/**
 * Samples for MetricAlerts List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/listMetricAlert.json
     */
    /**
     * Sample code: List metric alert rules.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void listMetricAlertRules(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getMetricAlerts().list(com.azure.core.util.Context.NONE);
    }
}
