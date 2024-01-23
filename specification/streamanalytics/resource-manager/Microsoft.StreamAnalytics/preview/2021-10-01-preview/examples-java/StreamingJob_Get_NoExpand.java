
/**
 * Samples for StreamingJobs GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
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

    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * StreamingJob_Get_Expand.json
     */
    /**
     * Sample code: Get a streaming job and use the $expand OData query parameter to expand inputs, outputs,
     * transformation, and functions.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void
        getAStreamingJobAndUseTheExpandODataQueryParameterToExpandInputsOutputsTransformationAndFunctions(
            com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.streamingJobs().getByResourceGroupWithResponse("sjrg3276", "sj7804",
            "inputs,outputs,transformation,functions", com.azure.core.util.Context.NONE);
    }
}
