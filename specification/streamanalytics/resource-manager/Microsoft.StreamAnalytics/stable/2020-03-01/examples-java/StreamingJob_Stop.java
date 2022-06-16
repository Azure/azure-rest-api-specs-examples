import com.azure.core.util.Context;

/** Samples for StreamingJobs Stop. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/StreamingJob_Stop.json
     */
    /**
     * Sample code: Stop a streaming job.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void stopAStreamingJob(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.streamingJobs().stop("sjrg6936", "sj59", Context.NONE);
    }
}
