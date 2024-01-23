
import com.azure.resourcemanager.streamanalytics.models.StreamingJob;

/**
 * Samples for StreamingJobs Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * StreamingJob_Update.json
     */
    /**
     * Sample code: Update a streaming job.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void updateAStreamingJob(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        StreamingJob resource = manager.streamingJobs()
            .getByResourceGroupWithResponse("sjrg6936", "sj59", null, com.azure.core.util.Context.NONE).getValue();
        resource.update().withEventsOutOfOrderMaxDelayInSeconds(21).withEventsLateArrivalMaxDelayInSeconds(13).apply();
    }
}
