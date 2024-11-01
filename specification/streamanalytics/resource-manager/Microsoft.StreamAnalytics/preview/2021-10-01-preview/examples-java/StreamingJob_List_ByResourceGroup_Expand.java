
/**
 * Samples for StreamingJobs ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * StreamingJob_List_ByResourceGroup_Expand.json
     */
    /**
     * Sample code: List all streaming jobs in a resource group and use the $expand OData query parameter to expand
     * inputs, outputs, transformation, and functions.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void
        listAllStreamingJobsInAResourceGroupAndUseTheExpandODataQueryParameterToExpandInputsOutputsTransformationAndFunctions(
            com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.streamingJobs().listByResourceGroup("sjrg3276", "inputs,outputs,transformation,functions",
            com.azure.core.util.Context.NONE);
    }
}
