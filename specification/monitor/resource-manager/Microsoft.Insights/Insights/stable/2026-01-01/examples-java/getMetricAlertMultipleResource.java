
/**
 * Samples for MetricAlerts GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/getMetricAlertMultipleResource.json
     */
    /**
     * Sample code: Get an alert rule for multiple resources.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void getAnAlertRuleForMultipleResources(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getMetricAlerts().getByResourceGroupWithResponse("gigtest",
            "MetricAlertOnMultipleResources", com.azure.core.util.Context.NONE);
    }
}
