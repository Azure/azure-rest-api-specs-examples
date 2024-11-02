
/**
 * Samples for Functions ListByStreamingJob.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Function_ListByStreamingJob.json
     */
    /**
     * Sample code: List all functions in a streaming job.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void
        listAllFunctionsInAStreamingJob(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.functions().listByStreamingJob("sjrg1637", "sj8653", null, com.azure.core.util.Context.NONE);
    }
}
