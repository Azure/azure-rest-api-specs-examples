
/**
 * Samples for MetricNamespaces List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-02-01/GetMetricNamespaces.json
     */
    /**
     * Sample code: Get Metric Namespaces without filter.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void getMetricNamespacesWithoutFilter(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getMetricNamespaces().list(
            "subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill",
            "2020-08-31T15:53:00Z", com.azure.core.util.Context.NONE);
    }
}
