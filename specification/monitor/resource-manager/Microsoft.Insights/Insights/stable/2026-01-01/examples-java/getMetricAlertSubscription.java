
/**
 * Samples for MetricAlerts GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/getMetricAlertSubscription.json
     */
    /**
     * Sample code: Get an alert rule on subscription.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void getAnAlertRuleOnSubscription(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getMetricAlerts().getByResourceGroupWithResponse("gigtest",
            "MetricAlertAtSubscriptionLevel", com.azure.core.util.Context.NONE);
    }
}
