
/**
 * Samples for MetricAlerts ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/listByResourceGroupMetricAlert.json
     */
    /**
     * Sample code: List metric alert rules.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void listMetricAlertRules(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getMetricAlerts().listByResourceGroup("gigtest", com.azure.core.util.Context.NONE);
    }
}
