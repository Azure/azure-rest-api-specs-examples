
/**
 * Samples for StreamingJobs List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/
     * StreamingJob_List_BySubscription_NoExpand.json
     */
    /**
     * Sample code: List all streaming jobs in a subscription and do not use the $expand OData query parameter.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void listAllStreamingJobsInASubscriptionAndDoNotUseTheExpandODataQueryParameter(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.streamingJobs().list(null, com.azure.core.util.Context.NONE);
    }
}
