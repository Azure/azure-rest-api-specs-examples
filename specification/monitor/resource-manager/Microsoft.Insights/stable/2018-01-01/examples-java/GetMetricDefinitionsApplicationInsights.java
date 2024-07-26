
/**
 * Samples for MetricDefinitions List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2018-01-01/examples/
     * GetMetricDefinitionsApplicationInsights.json
     */
    /**
     * Sample code: Get Application Insights Metric Definitions without filter.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getApplicationInsightsMetricDefinitionsWithoutFilter(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getMetricDefinitions().list(
            "subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill/providers/microsoft.insights/metricdefinitions",
            "microsoft.insights/components", com.azure.core.util.Context.NONE);
    }
}
