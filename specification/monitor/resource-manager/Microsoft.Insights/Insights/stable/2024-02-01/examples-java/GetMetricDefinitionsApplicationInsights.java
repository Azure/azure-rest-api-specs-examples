
/**
 * Samples for MetricDefinitions List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-02-01/GetMetricDefinitionsApplicationInsights.json
     */
    /**
     * Sample code: Get Application Insights Metric Definitions without filter.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void
        getApplicationInsightsMetricDefinitionsWithoutFilter(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getMetricDefinitions().list(
            "subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill/providers/microsoft.insights/metricdefinitions",
            "microsoft.insights/components", com.azure.core.util.Context.NONE);
    }
}
