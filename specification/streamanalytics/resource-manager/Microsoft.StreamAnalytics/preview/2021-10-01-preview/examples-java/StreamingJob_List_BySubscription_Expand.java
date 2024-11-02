
/**
 * Samples for StreamingJobs List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * StreamingJob_List_BySubscription_Expand.json
     */
    /**
     * Sample code: List all streaming jobs in a subscription and use the $expand OData query parameter to expand
     * inputs, outputs, transformation, and functions.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void
        listAllStreamingJobsInASubscriptionAndUseTheExpandODataQueryParameterToExpandInputsOutputsTransformationAndFunctions(
            com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.streamingJobs().list("inputs,outputs,transformation,functions", com.azure.core.util.Context.NONE);
    }
}
