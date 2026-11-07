
/**
 * Samples for MetricAlerts GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/getMetricAlertSingleResource.json
     */
    /**
     * Sample code: Get an alert rule for single resource.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void getAnAlertRuleForSingleResource(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getMetricAlerts().getByResourceGroupWithResponse("gigtest", "chiricutin",
            com.azure.core.util.Context.NONE);
    }
}
