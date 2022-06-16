import com.azure.core.util.Context;
import com.azure.resourcemanager.streamanalytics.models.ScaleStreamingJobParameters;

/** Samples for StreamingJobs Scale. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/StreamingJob_Scale.json
     */
    /**
     * Sample code: Scale a streaming job.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void scaleAStreamingJob(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager
            .streamingJobs()
            .scale("sjrg6936", "sj59", new ScaleStreamingJobParameters().withStreamingUnits(36), Context.NONE);
    }
}
