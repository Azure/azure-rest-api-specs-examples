
/**
 * Samples for MetricAlerts GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/getDynamicMetricAlertMultipleResource.json
     */
    /**
     * Sample code: Get a dynamic alert rule for multiple resources.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void
        getADynamicAlertRuleForMultipleResources(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getMetricAlerts().getByResourceGroupWithResponse("gigtest",
            "MetricAlertOnMultipleResources", com.azure.core.util.Context.NONE);
    }
}
