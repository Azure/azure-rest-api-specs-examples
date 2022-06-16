import com.azure.core.util.Context;

/** Samples for Inputs ListByStreamingJob. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Input_ListByStreamingJob_Diagnostics.json
     */
    /**
     * Sample code: List all inputs in a streaming job and include diagnostic information using the $select OData query
     * parameter.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void listAllInputsInAStreamingJobAndIncludeDiagnosticInformationUsingTheSelectODataQueryParameter(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.inputs().listByStreamingJob("sjrg3276", "sj7804", "*", Context.NONE);
    }
}
