import com.azure.core.util.Context;

/** Samples for Outputs ListByStreamingJob. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Output_ListByStreamingJob.json
     */
    /**
     * Sample code: List all outputs in a streaming job.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void listAllOutputsInAStreamingJob(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.outputs().listByStreamingJob("sjrg2157", "sj6458", null, Context.NONE);
    }
}
