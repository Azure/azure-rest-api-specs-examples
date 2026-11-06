
/**
 * Samples for MetricAlerts GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/getDynamicMetricAlertSingleResource.json
     */
    /**
     * Sample code: Get a dynamic alert rule for single resource.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void getADynamicAlertRuleForSingleResource(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getMetricAlerts().getByResourceGroupWithResponse("gigtest", "chiricutin",
            com.azure.core.util.Context.NONE);
    }
}
