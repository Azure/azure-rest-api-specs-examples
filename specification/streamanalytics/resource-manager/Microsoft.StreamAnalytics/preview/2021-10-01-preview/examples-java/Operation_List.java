
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Operation_List.json
     */
    /**
     * Sample code: List available operations for the Stream Analytics resource provider.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void listAvailableOperationsForTheStreamAnalyticsResourceProvider(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
