import com.azure.core.util.Context;

/** Samples for Functions ListByStreamingJob. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Function_ListByStreamingJob.json
     */
    /**
     * Sample code: List all functions in a streaming job.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void listAllFunctionsInAStreamingJob(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.functions().listByStreamingJob("sjrg1637", "sj8653", null, Context.NONE);
    }
}
