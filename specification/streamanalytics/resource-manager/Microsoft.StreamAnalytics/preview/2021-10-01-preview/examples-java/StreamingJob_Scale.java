
import com.azure.resourcemanager.streamanalytics.models.ScaleStreamingJobParameters;

/**
 * Samples for StreamingJobs Scale.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * StreamingJob_Scale.json
     */
    /**
     * Sample code: Scale a streaming job.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void scaleAStreamingJob(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.streamingJobs().scale("sjrg", "sjName", new ScaleStreamingJobParameters().withStreamingUnits(36),
            com.azure.core.util.Context.NONE);
    }
}
