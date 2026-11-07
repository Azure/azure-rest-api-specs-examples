
/**
 * Samples for MetricDefinitions ListAtSubscriptionScope.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-02-01/GetMultiResourceMetricDefinitions.json
     */
    /**
     * Sample code: Get Subscription level Metric Definitions without filter.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void
        getSubscriptionLevelMetricDefinitionsWithoutFilter(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getMetricDefinitions().listAtSubscriptionScope("westus2",
            "microsoft.compute/virtualmachines", com.azure.core.util.Context.NONE);
    }
}
