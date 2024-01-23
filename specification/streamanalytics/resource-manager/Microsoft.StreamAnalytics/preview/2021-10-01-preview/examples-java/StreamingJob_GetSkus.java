
/**
 * Samples for Sku List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * StreamingJob_GetSkus.json
     */
    /**
     * Sample code: Get valid SKUs list for the specified streaming job.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void getValidSKUsListForTheSpecifiedStreamingJob(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.skus().list("sjrg3276", "sj7804", com.azure.core.util.Context.NONE);
    }
}
