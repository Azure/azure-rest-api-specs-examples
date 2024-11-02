
/**
 * Samples for Outputs ListByStreamingJob.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Output_ListByStreamingJob.json
     */
    /**
     * Sample code: List all outputs in a streaming job.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void
        listAllOutputsInAStreamingJob(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.outputs().listByStreamingJob("sjrg2157", "sj6458", null, com.azure.core.util.Context.NONE);
    }
}
