
/**
 * Samples for MetricDefinitions List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-02-01/GetMetricDefinitions.json
     */
    /**
     * Sample code: Get Metric Definitions without filter.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void getMetricDefinitionsWithoutFilter(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getMetricDefinitions().list(
            "subscriptions/07c0b09d-9f69-4e6e-8d05-f59f67299cb2/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/alertruleTest/providers/microsoft.insights/metricDefinitions",
            "Microsoft.Web/sites", com.azure.core.util.Context.NONE);
    }
}
