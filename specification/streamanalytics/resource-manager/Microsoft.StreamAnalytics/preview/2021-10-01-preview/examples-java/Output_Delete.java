
/**
 * Samples for Outputs Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Output_Delete.json
     */
    /**
     * Sample code: Delete an output.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void deleteAnOutput(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.outputs().deleteWithResponse("sjrg2157", "sj6458", "output1755", com.azure.core.util.Context.NONE);
    }
}
