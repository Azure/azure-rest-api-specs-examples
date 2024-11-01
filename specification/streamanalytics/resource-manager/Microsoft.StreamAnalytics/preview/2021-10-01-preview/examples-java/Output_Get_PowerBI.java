
/**
 * Samples for Outputs Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Output_Get_PowerBI.json
     */
    /**
     * Sample code: Get a Power BI output.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void getAPowerBIOutput(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.outputs().getWithResponse("sjrg7983", "sj2331", "output3022", com.azure.core.util.Context.NONE);
    }
}
