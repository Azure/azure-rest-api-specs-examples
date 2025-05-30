
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/stable/2020-05-15/examples/
     * Operation_List.json
     */
    /**
     * Sample code: List available operations for the Time Series Insights resource provider.
     * 
     * @param manager Entry point to TimeSeriesInsightsManager.
     */
    public static void listAvailableOperationsForTheTimeSeriesInsightsResourceProvider(
        com.azure.resourcemanager.timeseriesinsights.TimeSeriesInsightsManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
