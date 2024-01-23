
/**
 * Samples for StreamingJobs Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * StreamingJob_Delete.json
     */
    /**
     * Sample code: Delete a streaming job.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void deleteAStreamingJob(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.streamingJobs().delete("sjrg6936", "sj59", com.azure.core.util.Context.NONE);
    }
}
