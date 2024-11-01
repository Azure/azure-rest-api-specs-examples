
/**
 * Samples for StreamingJobs ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * StreamingJob_List_ByResourceGroup_NoExpand.json
     */
    /**
     * Sample code: List all streaming jobs in a resource group and do not use the $expand OData query parameter.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void listAllStreamingJobsInAResourceGroupAndDoNotUseTheExpandODataQueryParameter(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.streamingJobs().listByResourceGroup("sjrg6936", null, com.azure.core.util.Context.NONE);
    }
}
