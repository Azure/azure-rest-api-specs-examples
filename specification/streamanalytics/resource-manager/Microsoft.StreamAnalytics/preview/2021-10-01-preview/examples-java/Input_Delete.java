
/**
 * Samples for Inputs Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Input_Delete.json
     */
    /**
     * Sample code: Delete an input.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void deleteAnInput(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.inputs().deleteWithResponse("sjrg8440", "sj9597", "input7225", com.azure.core.util.Context.NONE);
    }
}
