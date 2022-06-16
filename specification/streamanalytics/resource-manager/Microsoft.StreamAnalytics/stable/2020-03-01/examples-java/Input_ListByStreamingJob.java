import com.azure.core.util.Context;

/** Samples for Inputs ListByStreamingJob. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Input_ListByStreamingJob.json
     */
    /**
     * Sample code: List all inputs in a streaming job.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void listAllInputsInAStreamingJob(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.inputs().listByStreamingJob("sjrg8440", "sj9597", null, Context.NONE);
    }
}
