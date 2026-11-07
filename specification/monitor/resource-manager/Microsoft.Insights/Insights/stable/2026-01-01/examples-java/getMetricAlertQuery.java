
/**
 * Samples for MetricAlerts GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/getMetricAlertQuery.json
     */
    /**
     * Sample code: Get a query based alert rule.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void getAQueryBasedAlertRule(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getMetricAlerts().getByResourceGroupWithResponse("gigtest", "chiricutin",
            com.azure.core.util.Context.NONE);
    }
}
