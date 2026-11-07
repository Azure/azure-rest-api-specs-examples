
/**
 * Samples for MetricAlertsStatus List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/getMetricAlertStatus.json
     */
    /**
     * Sample code: Get an alert rule status.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void getAnAlertRuleStatus(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getMetricAlertsStatus().list("gigtest", "chiricutin", com.azure.core.util.Context.NONE);
    }
}
