
/**
 * Samples for Outputs Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Output_Get_DeltaLake.json
     */
    /**
     * Sample code: Get a Delta Lake output.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void getADeltaLakeOutput(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.outputs().getWithResponse("sjrg", "sjName", "output1221", com.azure.core.util.Context.NONE);
    }
}
