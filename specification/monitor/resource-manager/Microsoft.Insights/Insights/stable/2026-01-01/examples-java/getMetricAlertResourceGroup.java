
/**
 * Samples for MetricAlerts GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/getMetricAlertResourceGroup.json
     */
    /**
     * Sample code: Get an alert rule on resource group(s).
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void getAnAlertRuleOnResourceGroupS(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getMetricAlerts().getByResourceGroupWithResponse("gigtest1",
            "MetricAlertAtResourceGroupLevel", com.azure.core.util.Context.NONE);
    }
}
