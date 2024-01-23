
import com.azure.resourcemanager.streamanalytics.models.OutputStartMode;
import com.azure.resourcemanager.streamanalytics.models.StartStreamingJobParameters;

/**
 * Samples for StreamingJobs Start.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * StreamingJob_Start_LastOutputEventTime.json
     */
    /**
     * Sample code: Start a streaming job with LastOutputEventTime output start mode.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void startAStreamingJobWithLastOutputEventTimeOutputStartMode(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.streamingJobs().start("sjrg6936", "sj59",
            new StartStreamingJobParameters().withOutputStartMode(OutputStartMode.LAST_OUTPUT_EVENT_TIME),
            com.azure.core.util.Context.NONE);
    }
}
