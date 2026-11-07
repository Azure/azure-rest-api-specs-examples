
/**
 * Samples for MetricAlerts GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/getWebTestMetricAlert.json
     */
    /**
     * Sample code: Get a web test alert rule.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void getAWebTestAlertRule(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getMetricAlerts().getByResourceGroupWithResponse("rg-example", "webtest-name-example",
            com.azure.core.util.Context.NONE);
    }
}
