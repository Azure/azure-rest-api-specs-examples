
/**
 * Samples for Outputs Test.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Output_Test.
     * json
     */
    /**
     * Sample code: Test the connection for an output.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void
        testTheConnectionForAnOutput(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.outputs().test("sjrg2157", "sj6458", "output1755", null, com.azure.core.util.Context.NONE);
    }
}
