import com.azure.core.util.Context;

/** Samples for Outputs Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Output_Get_PowerBI.json
     */
    /**
     * Sample code: Get a Power BI output.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void getAPowerBIOutput(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.outputs().getWithResponse("sjrg7983", "sj2331", "output3022", Context.NONE);
    }
}
