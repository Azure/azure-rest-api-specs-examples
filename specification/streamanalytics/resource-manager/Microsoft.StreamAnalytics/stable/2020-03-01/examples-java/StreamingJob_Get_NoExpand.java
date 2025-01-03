
/**
 * Samples for StreamingJobs GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/
     * StreamingJob_Get_NoExpand.json
     */
    /**
     * Sample code: Get a streaming job and do not use the $expand OData query parameter.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void getAStreamingJobAndDoNotUseTheExpandODataQueryParameter(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.streamingJobs().getByResourceGroupWithResponse("sjrg6936", "sj59", null,
            com.azure.core.util.Context.NONE);
    }
}
