
/**
 * Samples for Outputs Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Output_Get_DocumentDB.json
     */
    /**
     * Sample code: Get a DocumentDB output.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void getADocumentDBOutput(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.outputs().getWithResponse("sjrg7983", "sj2331", "output3022", com.azure.core.util.Context.NONE);
    }
}
