
import com.azure.resourcemanager.streamanalytics.models.OutputStartMode;
import com.azure.resourcemanager.streamanalytics.models.StartStreamingJobParameters;

/**
 * Samples for StreamingJobs Start.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * StreamingJob_Start_JobStartTime.json
     */
    /**
     * Sample code: Start a streaming job with JobStartTime output start mode.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void startAStreamingJobWithJobStartTimeOutputStartMode(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.streamingJobs().start("sjrg6936", "sj59",
            new StartStreamingJobParameters().withOutputStartMode(OutputStartMode.JOB_START_TIME),
            com.azure.core.util.Context.NONE);
    }
}
