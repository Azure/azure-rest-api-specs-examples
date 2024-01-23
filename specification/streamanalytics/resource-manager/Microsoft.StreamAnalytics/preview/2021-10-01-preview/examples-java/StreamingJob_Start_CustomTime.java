
import com.azure.resourcemanager.streamanalytics.models.OutputStartMode;
import com.azure.resourcemanager.streamanalytics.models.StartStreamingJobParameters;
import java.time.OffsetDateTime;

/**
 * Samples for StreamingJobs Start.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * StreamingJob_Start_CustomTime.json
     */
    /**
     * Sample code: Start a streaming job with CustomTime output start mode.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void startAStreamingJobWithCustomTimeOutputStartMode(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.streamingJobs()
            .start(
                "sjrg6936", "sj59", new StartStreamingJobParameters().withOutputStartMode(OutputStartMode.CUSTOM_TIME)
                    .withOutputStartTime(OffsetDateTime.parse("2012-12-12T12:12:12Z")),
                com.azure.core.util.Context.NONE);
    }
}
