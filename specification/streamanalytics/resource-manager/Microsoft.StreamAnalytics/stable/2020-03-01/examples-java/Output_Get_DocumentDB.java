import com.azure.core.util.Context;

/** Samples for Outputs Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Output_Get_DocumentDB.json
     */
    /**
     * Sample code: Get a DocumentDB output.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void getADocumentDBOutput(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.outputs().getWithResponse("sjrg7983", "sj2331", "output3022", Context.NONE);
    }
}
